package controller

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/lyonnee/go-template/internal/application"
	"github.com/lyonnee/go-template/internal/interfaces/http/dto"
	"github.com/lyonnee/go-template/pkg/log"
)

type AuthController struct {
}

func NewAuthController() *AuthController {
	return &AuthController{}
}

func (c *AuthController) SignUp(ctx context.Context, reqCtx *app.RequestContext) {}

func (c *AuthController) Login(ctx context.Context, reqCtx *app.RequestContext) {
	var req dto.LoginReq

	// get params from body
	if err := reqCtx.Bind(&req); err != nil {
		dto.Fail(reqCtx, dto.CODE_INVALID_BODY_ARGUMENT, "")
		return
	}

	// validate params

	// generate cmd
	cmd := application.LoginCmd{
		PhoneNumber: req.PhoneNumber,
		Email:       req.Email,
		Password:    req.Password,
	}

	// execute cmd
	_, err := application.Login(ctx, &cmd)
	if err != nil {
		dto.Fail(reqCtx, dto.CODE_SERVER_ERROR, "")
		return
	}

	dto.Ok[any](reqCtx, "", "")
}

type S struct {
	str string
}

func (c *AuthController) RefreshToken(ctx context.Context, reqCtx *app.RequestContext) {
	var s *S
	log.Error(s.str)
}
