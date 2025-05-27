package controller

import (
	"context"
	"errors"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/lyonnee/go-template/internal/application/command_executor"
	"github.com/lyonnee/go-template/internal/application/query_executor"
	domainErrors "github.com/lyonnee/go-template/internal/domain/errors"
	"github.com/lyonnee/go-template/internal/infrastructure/log"
	"github.com/lyonnee/go-template/internal/interfaces/http/dto"
	"github.com/lyonnee/go-template/pkg/auth"
)

type UserController struct {
	userCmdService   *command_executor.UserCommandService
	userQueryService *query_executor.UserQueryService
	logger           log.Logger
}

func NewUserController(
	userCmdService *command_executor.UserCommandService,
	userQueryService *query_executor.UserQueryService,
	logger log.Logger) *UserController {
	return &UserController{
		userCmdService:   userCmdService,
		userQueryService: userQueryService,
		logger:           logger,
	}
}

// GetUser 获取用户信息
func (c *UserController) GetUser(ctx context.Context, reqCtx *app.RequestContext) {
	// 从路径参数获取用户ID
	userIDStr := reqCtx.Param("id")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		c.logger.ErrorKV("GetUser invalid user ID format", "userIdStr", userIDStr, "error", err)
		dto.Fail(reqCtx, dto.CODE_INVALID_PATH_ARGUMENT, "用户ID格式错误")
		return
	}

	c.logger.DebugKV("GetUser request received", "userId", userID)

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
		c.logger.WarnKV("GetUser unauthorized access attempt",
			"requestedUserId", userID,
			"authenticatedUserId", userClaims.UserId)
		dto.Fail(reqCtx, dto.CODE_TOKEN_INVALID, "无权查看该用户信息")
		return
	}

	// 获取用户信息
	user, err := c.userQueryService.GetUserById(ctx, userID)
	if err != nil {
		c.logger.ErrorKV("GetUser failed", "error", err, "userId", userID)
		if errors.Is(err, domainErrors.ErrUserNotFound) {
			dto.Fail(reqCtx, dto.CODE_INVALID_PATH_ARGUMENT, "用户不存在")
		} else {
			dto.Fail(reqCtx, dto.CODE_SERVER_ERROR, "获取用户信息失败")
		}
		return
	}

	c.logger.InfoKV("User information retrieved successfully", "userId", userID)

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
		c.logger.ErrorKV("UpdateUsername invalid user ID format", "userIdStr", userIDStr, "error", err)
		dto.Fail(reqCtx, dto.CODE_INVALID_PATH_ARGUMENT, "用户ID格式错误")
		return
	}

	c.logger.DebugKV("UpdateUsername request received", "userId", userID)

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
		c.logger.WarnKV("UpdateUsername unauthorized access attempt",
			"requestedUserId", userID,
			"authenticatedUserId", userClaims.UserId)
		dto.Fail(reqCtx, dto.CODE_TOKEN_INVALID, "无权修改该用户信息")
		return
	}

	// 绑定参数
	var req dto.UpdateUsernameReq
	if err := reqCtx.Bind(&req); err != nil {
		c.logger.ErrorKV("UpdateUsername bind params failed", "error", err, "userId", userID)
		dto.Fail(reqCtx, dto.CODE_INVALID_BODY_ARGUMENT, "参数格式错误")
		return
	}

	c.logger.DebugKV("UpdateUsername request bound successfully",
		"userId", userID,
		"newUsername", req.Username)

	// 创建命令
	cmd := &command_executor.UpdateUsernameCmd{
		UserID:   userID,
		Username: req.Username,
	}

	// 执行更新
	user, err := c.userCmdService.UpdateUsername(ctx, cmd)
	if err != nil {
		c.logger.ErrorKV("UpdateUsername failed", "error", err, "userId", userID, "newUsername", req.Username)

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

	c.logger.InfoKV("Username updated successfully",
		"userId", userID,
		"newUsername", req.Username)

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
