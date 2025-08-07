package controller

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/lyonnee/go-template/internal/application/service"
	"github.com/lyonnee/go-template/internal/interfaces/http/dto"
	"github.com/lyonnee/go-template/pkg/di"
	"github.com/lyonnee/go-template/pkg/log"
	"go.uber.org/zap"
)

type AuthController struct {
	authCmdService *service.AuthCommandService
	logger         *log.Logger
}

func init() {
	di.AddSingleton[*AuthController](NewAuthController)
}

func NewAuthController() (*AuthController, error) {
	return &AuthController{
		authCmdService: di.Get[*service.AuthCommandService](),
		logger:         di.Get[*log.Logger](),
	}, nil
}

// Login 用户登录
func (c *AuthController) Login(ctx context.Context, reqCtx *app.RequestContext) {
	c.logger.Debug("Login request received")

	var req dto.LoginReq

	// 绑定参数
	if err := reqCtx.Bind(&req); err != nil {
		c.logger.Error("Login bind params failed", zap.Error(err))
		dto.Fail(reqCtx, dto.CODE_INVALID_BODY_ARGUMENT, "参数格式错误")
		return
	}

	c.logger.Debug("Login request bound successfully", zap.String("username", req.Username))

	// 创建命令
	cmd := &service.LoginCmd{
		Username: req.Username,
		Password: req.Password,
	}

	// 执行登录
	result, err := c.authCmdService.Login(ctx, cmd)
	if err != nil {
		c.logger.Error("Login failed", zap.Error(err), zap.String("username", req.Username))
		dto.Fail(reqCtx, dto.CODE_INVALID_BODY_ARGUMENT, "用户名或密码错误")
		return
	}

	c.logger.Info("User logged in successfully", zap.String("username", req.Username))

	// 构造响应
	resp := dto.LoginResp{
		AccessToken:  result.AccessToken,
		RefreshToken: result.RefreshToken,
	}

	dto.Ok(reqCtx, "登录成功", resp)
}

// RefreshToken 刷新token
func (c *AuthController) RefreshToken(ctx context.Context, reqCtx *app.RequestContext) {
	c.logger.Debug("RefreshToken request received")

	var req dto.RefreshTokenReq

	// 绑定参数
	if err := reqCtx.Bind(&req); err != nil {
		c.logger.Error("RefreshToken bind params failed", zap.Error(err))
		dto.Fail(reqCtx, dto.CODE_INVALID_BODY_ARGUMENT, "参数格式错误")
		return
	}

	c.logger.Debug("RefreshToken request bound successfully")

	// 创建命令
	cmd := &service.RefreshTokenCmd{
		RefreshToken: req.RefreshToken,
	}

	// 执行刷新
	result, err := c.authCmdService.RefreshToken(ctx, cmd)
	if err != nil {
		c.logger.Error("RefreshToken failed", zap.Error(err))
		dto.Fail(reqCtx, dto.CODE_TOKEN_INVALID, "刷新token无效")
		return
	}

	c.logger.Info("Token refreshed successfully")

	// 构造响应
	resp := dto.RefreshTokenResp{
		AccessToken: result.AccessToken,
	}

	dto.Ok(reqCtx, "刷新成功", resp)
}
