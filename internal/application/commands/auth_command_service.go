package commands

import (
	"context"
	"errors"

	"github.com/lyonnee/go-template/pkg/log"

	"github.com/lyonnee/go-template/internal/domain/entity"
	"github.com/lyonnee/go-template/internal/domain/repository"
	"github.com/lyonnee/go-template/internal/infrastructure/auth"
	"github.com/lyonnee/go-template/internal/infrastructure/database"
	"github.com/lyonnee/go-template/pkg/di"
	"go.uber.org/zap"
)

type AuthCommandService struct {
	logger    *log.Logger
	dbContext *database.Database

	userRepo repository.UserRepository
}

func init() {
	di.AddSingleton[*AuthCommandService](NewAuthCommandService)
}

// NewAuthService 创建认证服务
func NewAuthCommandService() (*AuthCommandService, error) {
	return &AuthCommandService{
		logger:    di.Get[*log.Logger](),
		dbContext: di.Get[*database.Database](),

		userRepo: di.Get[repository.UserRepository](),
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
	if err := s.dbContext.Conn(ctx, func(ctx context.Context) error {
		// 查找用户
		userInfo, err := s.userRepo.FindByUsername(ctx, cmd.Username)
		if err != nil {
			return err
		}

		user = userInfo

		if err := user.Login(cmd.Password); err != nil {
			s.logger.Warn("Login failed - invalid password", zap.String("username", cmd.Username), zap.Uint64("userId", user.ID))
			return nil, errors.New("invalid username or password")
		}

		jwtGenerator := di.Get[*auth.JWTGenerator]()

		accessToken, err := jwtGenerator.GenerateAccessToken(user.ID, user.Username)
		if err != nil {
			return nil, err
		}

		refreshToken, err := jwtGenerator.GenerateRefreshToken(user.ID, user.Username)
		if err != nil {
			return nil, err
		}

		return nil
	}); err != nil {
		s.logger.Error("Database connection failed", zap.Error(err))
		return nil, err
	}

	s.logger.Info("User logged in successfully", zap.String("username", cmd.Username), zap.Uint64("userId", user.ID))

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
	jwtManager := di.Get[*auth.JWTGenerator]()

	// 验证刷新token
	claims, err := jwtManager.ValidateToken(cmd.RefreshToken)
	if err != nil {
		s.logger.Warn("Invalid refresh token provided", zap.Error(err))
		return nil, errors.New("invalid refresh token")
	}

	// 生成新的访问token
	newAccessToken, err := jwtManager.GenerateAccessToken(claims.UserId, claims.AlternativeID)
	if err != nil {
		s.logger.Error("Failed to generate new access token", zap.Error(err), zap.Uint64("userId", claims.UserId))
		return nil, err
	}

	s.logger.Info("Access token refreshed successfully", zap.Uint64("userId", claims.UserId))

	return &RefreshTokenResult{
		AccessToken: newAccessToken,
	}, nil
}
