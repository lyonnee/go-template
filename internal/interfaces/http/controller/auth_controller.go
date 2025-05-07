package controller

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/lyonnee/go-template/internal/application"
	"github.com/lyonnee/go-template/internal/interfaces/http/dto"
)

type AuthController struct {
	authService application.AuthService
	log.Logger
}

func NewAuthController(
	authService application.AuthService,
) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (c *AuthController) Login(ctx *fiber.Ctx) error {
	var req dto.LoginReq

	// get params
	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	// validate params

	// generate cmd
	cmd := application.LoginCmd{
		PhoneNumber: req.PhoneNumber,
		Email:       req.Email,
		Password:    req.Password,
	}

	// execute cmd
	_, err := c.authService.Login(ctx.UserContext(), &cmd)
	if err != nil {
		return err
	}

	return nil
}

func (c *AuthController) RefreshToken(ctx *fiber.Ctx) error {
	return nil
}
