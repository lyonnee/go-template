package controller

import (
	"context"
	"errors"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/lyonnee/go-template/bootstrap/di"
	"github.com/lyonnee/go-template/internal/application/service"
	domainErrors "github.com/lyonnee/go-template/internal/domain/errors"
	"github.com/lyonnee/go-template/internal/interfaces/http/dto"
	"go.uber.org/zap"
)

type AuthController struct {
	authCmdService *service.AuthCommandService
	logger         *zap.Logger
}

func NewAuthController() (*AuthController, error) {
	return &AuthController{
		authCmdService: di.Get[*service.AuthCommandService](),
		logger:         di.Get[*zap.Logger](),
	}, nil
}

// SignUp 用户注册
func (c *AuthController) SignUp(ctx context.Context, reqCtx *app.RequestContext) {
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
	result, err := c.authCmdService.SignUp(ctx, cmd)
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
