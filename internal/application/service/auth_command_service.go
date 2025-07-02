package service

import (
	"context"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/lyonnee/go-template/bootstrap/di"
	"github.com/lyonnee/go-template/internal/domain/entity"
	"github.com/lyonnee/go-template/internal/domain/repository"
	"github.com/lyonnee/go-template/internal/domain/service"
	"github.com/lyonnee/go-template/internal/infrastructure/auth"
	"github.com/lyonnee/go-template/internal/infrastructure/persistence"
	"go.uber.org/zap"
)

type AuthCommandService struct {
	logger    *zap.Logger
	dbContext persistence.DBContext

	userRepo repository.UserRepository
}

// NewAuthService 创建认证服务
func NewAuthCommandService() (*AuthCommandService, error) {
	return &AuthCommandService{
		logger:    di.Get[*zap.Logger](),
		dbContext: di.Get[persistence.DBContext](),
		userRepo:  di.Get[repository.UserRepository](),
	}, nil
}

// SignUpCmd 注册命令
type SignUpCmd struct {
	Username string
	Password string
	Email    string
	Phone    string
}

// SignUpResult 注册结果
type SignUpResult struct {
	AccessToken  string
	RefreshToken string
	User         *entity.User
}

// SignUp 用户注册
func (s *AuthCommandService) SignUp(ctx context.Context, cmd *SignUpCmd) (*SignUpResult, error) {
	s.logger.Info("Starting user registration",
		zap.String("username", cmd.Username),
		zap.String("email", cmd.Email))

	var user *entity.User
	if err := s.dbContext.Conn(func(c *sqlx.Conn) error {
		userDomainService := service.NewUserService(s.userRepo, s.logger)
		newUser, err := userDomainService.CreateUser(ctx, cmd.Username, cmd.Password, cmd.Email, cmd.Phone)
		if err != nil {
			s.logger.Error("Failed to create user", zap.Error(err))
			return err
		}

		if err := s.userRepo.Create(ctx, newUser); err != nil {
			s.logger.Error("Failed to create user in database", zap.Error(err))
			return err
		}

		return nil
	}); err != nil {
		s.logger.Error("User registration failed", zap.Error(err), zap.String("username", cmd.Username))
		return nil, err
	}

	jwtManager := di.Get[*auth.JWTManager]()
	// 生成token
	accessToken, err := jwtManager.GenerateAccessToken(user.ID, user.Username)
	if err != nil {
		s.logger.Error("Failed to generate access token for new user", zap.Error(err), zap.Int64("userId", user.ID))
		return nil, err
	}

	refreshToken, err := jwtManager.GenerateRefreshToken(user.ID, user.Username)
	if err != nil {
		s.logger.Error("Failed to generate refresh token for new user", zap.Error(err), zap.Int64("userId", user.ID))
		return nil, err
	}

	s.logger.Info("User registration completed successfully",
		zap.String("username", cmd.Username),
		zap.Int64("userId", user.ID))

	return &SignUpResult{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         user,
	}, nil
}

// LoginCmd 登录命令
type LoginCmd struct {
	Username string
	Password string
}

// LoginResult 登录结果
type LoginResult struct {
	AccessToken  string
	RefreshToken string
}

// Login 用户登录
func (s *AuthCommandService) Login(ctx context.Context, cmd *LoginCmd) (*LoginResult, error) {
	s.logger.Debug("Login attempt", zap.String("username", cmd.Username))

	var user *entity.User
	if err := s.dbContext.Conn(func(c *sqlx.Conn) error {
		// 查找用户
		userInfo, err := s.userRepo.FindByUsername(ctx, cmd.Username)
		if err != nil {
			return err
		}

		user = userInfo

		return nil
	}); err != nil {
		s.logger.Error("Database connection failed", zap.Error(err))
		return nil, err
	}

	accessToken, refreshToken, err := user.Login(cmd.Password)
	if err != nil {
		s.logger.Warn("Login failed - invalid password", zap.String("username", cmd.Username), zap.Int64("userId", user.ID))
		return nil, errors.New("invalid username or password")
	}

	s.logger.Info("User logged in successfully", zap.String("username", cmd.Username), zap.Int64("userId", user.ID))

	return &LoginResult{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// RefreshTokenCmd 刷新token命令
type RefreshTokenCmd struct {
	RefreshToken string
}

// RefreshTokenResult 刷新token结果
type RefreshTokenResult struct {
	AccessToken string
}

// RefreshToken 刷新token
func (s *AuthCommandService) RefreshToken(ctx context.Context, cmd *RefreshTokenCmd) (*RefreshTokenResult, error) {
	s.logger.Debug("RefreshToken called")

	jwtManager := di.Get[*auth.JWTManager]()
	// 验证刷新token
	claims, err := jwtManager.ValidateToken(cmd.RefreshToken)
	if err != nil {
		s.logger.Warn("Invalid refresh token provided", zap.Error(err))
		return nil, errors.New("invalid refresh token")
	}

	// 生成新的访问token
	newAccessToken, err := jwtManager.GenerateAccessToken(claims.UserId, claims.AlternativeID)
	if err != nil {
		s.logger.Error("Failed to generate new access token", zap.Error(err), zap.Int64("userId", claims.UserId))
		return nil, err
	}

	s.logger.Info("Access token refreshed successfully", zap.Int64("userId", claims.UserId))

	return &RefreshTokenResult{
		AccessToken: newAccessToken,
	}, nil
}
