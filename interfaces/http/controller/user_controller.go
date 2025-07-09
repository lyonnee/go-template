package controller

import (
	"context"
	"errors"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/lyonnee/go-template/application/service"
	"github.com/lyonnee/go-template/bootstrap/di"
	domainErrors "github.com/lyonnee/go-template/domain/errors"
	"github.com/lyonnee/go-template/infrastructure/auth"
	"github.com/lyonnee/go-template/interfaces/http/dto"
	"go.uber.org/zap"
)

type UserController struct {
	userCmdService   *service.UserCommandService
	userQueryService *service.UserQueryService
	logger           *zap.Logger
}

func NewUserController() (*UserController, error) {
	return &UserController{
		userCmdService:   di.Get[*service.UserCommandService](),
		userQueryService: di.Get[*service.UserQueryService](),
		logger:           di.Get[*zap.Logger](),
	}, nil
}

// Register 用户注册
func (c *UserController) Register(ctx context.Context, reqCtx *app.RequestContext) {
	c.logger.Debug("SignUp request received")

	var req dto.SignUpReq

	// 绑定参数
	if err := reqCtx.Bind(&req); err != nil {
		c.logger.Error("SignUp bind params failed", zap.Error(err))
		dto.Fail(reqCtx, dto.CODE_INVALID_BODY_ARGUMENT, "参数格式错误")
		return
	}

	c.logger.Debug("SignUp request bound successfully", zap.String("username", req.Username), zap.String("email", req.Email))

	// 创建命令
	cmd := &service.SignUpCmd{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
		Phone:    req.Phone,
	}

	// 执行注册
	result, err := c.userCmdService.SignUp(ctx, cmd)
	if err != nil {
		c.logger.Error("SignUp failed", zap.Error(err), zap.String("username", req.Username))

		// 处理业务错误
		switch {
		case errors.Is(err, domainErrors.ErrUsernameTaken):
			dto.Fail(reqCtx, dto.CODE_INVALID_BODY_ARGUMENT, "用户名已被使用")
		case errors.Is(err, domainErrors.ErrEmailTaken):
			dto.Fail(reqCtx, dto.CODE_INVALID_BODY_ARGUMENT, "邮箱已被使用")
		case errors.Is(err, domainErrors.ErrPhoneTaken):
			dto.Fail(reqCtx, dto.CODE_INVALID_BODY_ARGUMENT, "手机号已被使用")
		default:
			dto.Fail(reqCtx, dto.CODE_SERVER_ERROR, "注册失败")
		}
		return
	}

	c.logger.Info("User registered successfully", zap.String("username", req.Username), zap.Int64("userId", result.User.ID))

	// 构造响应
	resp := dto.SignUpResp{
		AccessToken:  result.AccessToken,
		RefreshToken: result.RefreshToken,
		User: &dto.UserInfo{
			ID:       result.User.ID,
			Username: result.User.Username,
			Email:    result.User.Email,
			Phone:    result.User.Phone,
		},
	}

	dto.Ok(reqCtx, "注册成功", resp)
}

// GetUser 获取用户信息
func (c *UserController) GetUser(ctx context.Context, reqCtx *app.RequestContext) {
	// 从路径参数获取用户ID
	userIDStr := reqCtx.Param("id")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		c.logger.Error("GetUser invalid user ID format", zap.String("userIdStr", userIDStr), zap.Error(err))
		dto.Fail(reqCtx, dto.CODE_INVALID_PATH_ARGUMENT, "用户ID格式错误")
		return
	}

	c.logger.Debug("GetUser request received", zap.Int64("userId", userID))

	// 获取当前登录用户信息
	claims, exists := reqCtx.Get("claims")
	if !exists {
		c.logger.Error("GetUser - no claims found in context")
		dto.Fail(reqCtx, dto.CODE_NOT_TOKEN, "未获取到用户信息")
		return
	}

	userClaims, ok := claims.(*auth.Claims)
	if !ok {
		c.logger.Error("GetUser - invalid claims type in context")
		dto.Fail(reqCtx, dto.CODE_TOKEN_INVALID, "用户信息格式错误")
		return
	}

	// 检查权限：只能查看自己的信息
	if userClaims.UserId != userID {
		c.logger.Warn("GetUser unauthorized access attempt",
			zap.Int64("requestedUserId", userID),
			zap.Int64("authenticatedUserId", userClaims.UserId))
		dto.Fail(reqCtx, dto.CODE_TOKEN_INVALID, "无权查看该用户信息")
		return
	}

	// 获取用户信息
	user, err := c.userQueryService.GetUserById(ctx, userID)
	if err != nil {
		c.logger.Error("GetUser failed", zap.Error(err), zap.Int64("userId", userID))
		if errors.Is(err, domainErrors.ErrUserNotFound) {
			dto.Fail(reqCtx, dto.CODE_INVALID_PATH_ARGUMENT, "用户不存在")
		} else {
			dto.Fail(reqCtx, dto.CODE_SERVER_ERROR, "获取用户信息失败")
		}
		return
	}

	c.logger.Info("User information retrieved successfully", zap.Int64("userId", userID))

	// 构造响应
	resp := dto.GetUserResp{
		User: &dto.UserInfo{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
			Phone:    user.Phone,
		},
	}

	dto.Ok(reqCtx, "获取成功", resp)
}

// UpdateUsername 修改用户名
func (c *UserController) UpdateUsername(ctx context.Context, reqCtx *app.RequestContext) {
	// 从路径参数获取用户ID
	userIDStr := reqCtx.Param("id")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		c.logger.Error("UpdateUsername invalid user ID format", zap.String("userIdStr", userIDStr), zap.Error(err))
		dto.Fail(reqCtx, dto.CODE_INVALID_PATH_ARGUMENT, "用户ID格式错误")
		return
	}

	c.logger.Debug("UpdateUsername request received", zap.Int64("userId", userID))

	// 获取当前登录用户信息
	claims, exists := reqCtx.Get("claims")
	if !exists {
		c.logger.Error("UpdateUsername - no claims found in context")
		dto.Fail(reqCtx, dto.CODE_NOT_TOKEN, "未获取到用户信息")
		return
	}

	userClaims, ok := claims.(*auth.Claims)
	if !ok {
		c.logger.Error("UpdateUsername - invalid claims type in context")
		dto.Fail(reqCtx, dto.CODE_TOKEN_INVALID, "用户信息格式错误")
		return
	}

	// 检查权限：只能修改自己的信息
	if userClaims.UserId != userID {
		c.logger.Warn("UpdateUsername unauthorized access attempt",
			zap.Int64("requestedUserId", userID),
			zap.Int64("authenticatedUserId", userClaims.UserId))
		dto.Fail(reqCtx, dto.CODE_TOKEN_INVALID, "无权修改该用户信息")
		return
	}

	// 绑定参数
	var req dto.UpdateUsernameReq
	if err := reqCtx.Bind(&req); err != nil {
		c.logger.Error("UpdateUsername bind params failed", zap.Error(err), zap.Int64("userId", userID))
		dto.Fail(reqCtx, dto.CODE_INVALID_BODY_ARGUMENT, "参数格式错误")
		return
	}

	c.logger.Debug("UpdateUsername request bound successfully",
		zap.Int64("userId", userID),
		zap.String("newUsername", req.Username))

	// 创建命令
	cmd := &service.UpdateUsernameCmd{
		UserID:   userID,
		Username: req.Username,
	}

	// 执行更新
	user, err := c.userCmdService.UpdateUsername(ctx, cmd)
	if err != nil {
		c.logger.Error("UpdateUsername failed", zap.Error(err), zap.Int64("userId", userID), zap.String("newUsername", req.Username))

		switch {
		case errors.Is(err, domainErrors.ErrUserNotFound):
			dto.Fail(reqCtx, dto.CODE_INVALID_PATH_ARGUMENT, "用户不存在")
		case errors.Is(err, domainErrors.ErrUsernameTaken):
			dto.Fail(reqCtx, dto.CODE_INVALID_BODY_ARGUMENT, "用户名已被使用")
		default:
			dto.Fail(reqCtx, dto.CODE_SERVER_ERROR, "修改用户名失败")
		}
		return
	}

	c.logger.Info("Username updated successfully", zap.Int64("userId", userID), zap.String("newUsername", req.Username))

	// 构造响应
	resp := dto.UpdateUsernameResp{
		User: &dto.UserInfo{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
			Phone:    user.Phone,
		},
	}

	dto.Ok(reqCtx, "修改成功", resp)
}
