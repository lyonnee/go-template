package controller

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lyonnee/go-template/internal/application/commands"
	"github.com/lyonnee/go-template/internal/application/queries"
	domainErrors "github.com/lyonnee/go-template/internal/domain/errors"
	"github.com/lyonnee/go-template/internal/infrastructure/auth"
	"github.com/lyonnee/go-template/internal/interfaces/http/dto"
	"github.com/lyonnee/go-template/pkg/di"
	"github.com/lyonnee/go-template/pkg/log"
	"go.uber.org/zap"
)

type UserController struct {
	userCmdService   *commands.UserCommandService
	userQueryService *queries.UserQueryService
	logger           *log.Logger
}

func init() {
	di.AddSingleton[*UserController](NewUserController)
}

func NewUserController() (*UserController, error) {
	return &UserController{
		userCmdService:   di.Get[*commands.UserCommandService](),
		userQueryService: di.Get[*queries.UserQueryService](),
		logger:           di.Get[*log.Logger](),
	}, nil
}

// Register 用户注册
func (c *UserController) Register(ctx *gin.Context) {
	c.logger.Debug("SignUp request received")

	var req dto.SignUpReq

	// 绑定参数
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.logger.Error("SignUp bind params failed", zap.Error(err))
		dto.Fail(ctx, dto.CODE_INVALID_BODY_ARGUMENT, "参数格式错误")
		return
	}

	c.logger.Debug("SignUp request bound successfully", zap.String("username", req.Username), zap.String("email", req.Email))

	// 创建命令
	cmd := &commands.SignUpCmd{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
		Phone:    req.Phone,
	}

	// 执行注册
	result, err := c.userCmdService.SignUp(ctx.Request.Context(), cmd)
	if err != nil {
		c.logger.Error("SignUp failed", zap.Error(err), zap.String("username", req.Username))

		// 处理业务错误
		switch {
		case errors.Is(err, domainErrors.ErrUsernameTaken):
			dto.Fail(ctx, dto.CODE_INVALID_BODY_ARGUMENT, "用户名已被使用")
		case errors.Is(err, domainErrors.ErrEmailTaken):
			dto.Fail(ctx, dto.CODE_INVALID_BODY_ARGUMENT, "邮箱已被使用")
		case errors.Is(err, domainErrors.ErrPhoneTaken):
			dto.Fail(ctx, dto.CODE_INVALID_BODY_ARGUMENT, "手机号已被使用")
		default:
			dto.Fail(ctx, dto.CODE_SERVER_ERROR, "注册失败")
		}
		return
	}

	c.logger.Info("User registered successfully", zap.String("username", req.Username), zap.Uint64("userId", result.User.ID))

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

	dto.Ok(ctx, "注册成功", resp)
}

// GetUser 获取用户信息
func (c *UserController) GetUser(ctx *gin.Context) {
	// 从路径参数获取用户ID
	userIDStr := ctx.Param("id")
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		c.logger.Error("GetUser invalid user ID format", zap.String("userIdStr", userIDStr), zap.Error(err))
		dto.Fail(ctx, dto.CODE_INVALID_PATH_ARGUMENT, "用户ID格式错误")
		return
	}

	c.logger.Debug("GetUser request received", zap.Uint64("userId", userID))

	// 获取当前登录用户信息
	claims, exists := ctx.Get("claims")
	if !exists {
		c.logger.Error("GetUser - no claims found in context")
		dto.Fail(ctx, dto.CODE_NOT_TOKEN, "未获取到用户信息")
		return
	}

	userClaims, ok := claims.(*auth.Claims)
	if !ok {
		c.logger.Error("GetUser - invalid claims type in context")
		dto.Fail(ctx, dto.CODE_TOKEN_INVALID, "用户信息格式错误")
		return
	}

	// 检查权限：只能查看自己的信息
	if userClaims.UserId != userID {
		c.logger.Warn("GetUser unauthorized access attempt",
			zap.Uint64("requestedUserId", userID),
			zap.Uint64("authenticatedUserId", userClaims.UserId))
		dto.Fail(ctx, dto.CODE_TOKEN_INVALID, "无权查看该用户信息")
		return
	}

	// 获取用户信息
	user, err := c.userQueryService.GetUserById(ctx.Request.Context(), userID)
	if err != nil {
		c.logger.Error("GetUser failed", zap.Error(err), zap.Uint64("userId", userID))
		if errors.Is(err, domainErrors.ErrUserNotFound) {
			dto.Fail(ctx, dto.CODE_INVALID_PATH_ARGUMENT, "用户不存在")
		} else {
			dto.Fail(ctx, dto.CODE_SERVER_ERROR, "获取用户信息失败")
		}
		return
	}

	c.logger.Info("User information retrieved successfully", zap.Uint64("userId", userID))

	// 构造响应
	resp := dto.GetUserResp{
		User: &dto.UserInfo{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
			Phone:    user.Phone,
		},
	}

	dto.Ok(ctx, "获取成功", resp)
}

// UpdateUsername 修改用户名
func (c *UserController) UpdateUsername(ctx *gin.Context) {
	// 从路径参数获取用户ID
	userIDStr := ctx.Param("id")
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		c.logger.Error("UpdateUsername invalid user ID format", zap.String("userIdStr", userIDStr), zap.Error(err))
		dto.Fail(ctx, dto.CODE_INVALID_PATH_ARGUMENT, "用户ID格式错误")
		return
	}

	c.logger.Debug("UpdateUsername request received", zap.Uint64("userId", userID))

	// 获取当前登录用户信息
	claims, exists := ctx.Get("claims")
	if !exists {
		c.logger.Error("UpdateUsername - no claims found in context")
		dto.Fail(ctx, dto.CODE_NOT_TOKEN, "未获取到用户信息")
		return
	}

	userClaims, ok := claims.(*auth.Claims)
	if !ok {
		c.logger.Error("UpdateUsername - invalid claims type in context")
		dto.Fail(ctx, dto.CODE_TOKEN_INVALID, "用户信息格式错误")
		return
	}

	// 检查权限：只能修改自己的信息
	if userClaims.UserId != userID {
		c.logger.Warn("UpdateUsername unauthorized access attempt",
			zap.Uint64("requestedUserId", userID),
			zap.Uint64("authenticatedUserId", userClaims.UserId))
		dto.Fail(ctx, dto.CODE_TOKEN_INVALID, "无权修改该用户信息")
		return
	}

	// 绑定参数
	var req dto.UpdateUsernameReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.logger.Error("UpdateUsername bind params failed", zap.Error(err), zap.Uint64("userId", userID))
		dto.Fail(ctx, dto.CODE_INVALID_BODY_ARGUMENT, "参数格式错误")
		return
	}

	c.logger.Debug("UpdateUsername request bound successfully",
		zap.Uint64("userId", userID),
		zap.String("newUsername", req.Username))

	// 创建命令
	cmd := &commands.UpdateUsernameCmd{
		UserID:   userID,
		Username: req.Username,
	}

	// 执行更新
	user, err := c.userCmdService.UpdateUsername(ctx.Request.Context(), cmd)
	if err != nil {
		c.logger.Error("UpdateUsername failed", zap.Error(err), zap.Uint64("userId", userID), zap.String("newUsername", req.Username))

		switch {
		case errors.Is(err, domainErrors.ErrUserNotFound):
			dto.Fail(ctx, dto.CODE_INVALID_PATH_ARGUMENT, "用户不存在")
		case errors.Is(err, domainErrors.ErrUsernameTaken):
			dto.Fail(ctx, dto.CODE_INVALID_BODY_ARGUMENT, "用户名已被使用")
		default:
			dto.Fail(ctx, dto.CODE_SERVER_ERROR, "修改用户名失败")
		}
		return
	}

	c.logger.Info("Username updated successfully", zap.Uint64("userId", userID), zap.String("newUsername", req.Username))

	// 构造响应
	resp := dto.UpdateUsernameResp{
		User: &dto.UserInfo{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
			Phone:    user.Phone,
		},
	}

	dto.Ok(ctx, "修改成功", resp)
}
