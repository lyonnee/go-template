package command_executor

import (
	"context"
	"errors"

	"github.com/lyonnee/go-template/internal/domain"
	"github.com/lyonnee/go-template/internal/domain/entity"
	domainErrors "github.com/lyonnee/go-template/internal/domain/errors"
	"github.com/lyonnee/go-template/internal/domain/repository"
	"github.com/lyonnee/go-template/internal/infrastructure/log"
	"github.com/lyonnee/go-template/pkg/auth"
	"github.com/lyonnee/go-template/pkg/persistence"
)

type AuthCommandService struct {
	userRepo repository.UserRepository
	logger   log.Logger
}

// NewAuthService 创建认证服务
func NewAuthCommandService(userRepo repository.UserRepository, logger log.Logger) *AuthCommandService {
	return &AuthCommandService{
		userRepo: userRepo,
		logger:   logger,
	}
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

// LoginCmd 登录命令
type LoginCmd struct {
	Username string
	Password string
}

// LoginResult 登录结果
type LoginResult struct {
	AccessToken  string
	RefreshToken string
	User         *entity.User
}

// RefreshTokenCmd 刷新token命令
type RefreshTokenCmd struct {
	RefreshToken string
}

// RefreshTokenResult 刷新token结果
type RefreshTokenResult struct {
	AccessToken string
}

// SignUp 用户注册
func (s *AuthCommandService) SignUp(ctx context.Context, cmd *SignUpCmd) (*SignUpResult, error) {
	s.logger.InfoKV("Starting user registration",
		"username", cmd.Username,
		"email", cmd.Email)

	// 开启事务
	tx, err := persistence.NewTx(ctx)
	if err != nil {
		s.logger.ErrorKV("Failed to start transaction", "error", err)
		return nil, err
	}
	defer tx.Rollback()
	userRepoWithTx := s.userRepo.WithExecutor(tx)

	userDomainService := domain.NewUserDomainService(userRepoWithTx, s.logger)
	user, err := userDomainService.CreateUser(ctx, cmd.Username, cmd.Password, cmd.Email, cmd.Phone)
	if err != nil {
		s.logger.ErrorKV("Failed to create user", "error", err)
		return nil, err
	}

	if err := userRepoWithTx.Create(ctx, user); err != nil {
		s.logger.ErrorKV("Failed to create user in database", "error", err)
		return nil, err
	}

	// 生成token
	accessToken, err := auth.GenerateAccessToken(user.ID, user.Username)
	if err != nil {
		s.logger.ErrorKV("Failed to generate access token for new user",
			"error", err,
			"userId", user.ID)
		return nil, err
	}

	refreshToken, err := auth.GenerateRefreshToken(user.ID, user.Username)
	if err != nil {
		s.logger.ErrorKV("Failed to generate refresh token for new user",
			"error", err,
			"userId", user.ID)
		return nil, err
	}

	// 提交事务
	if err := tx.Commit(); err != nil {
		s.logger.ErrorKV("Failed to commit user registration transaction",
			"error", err,
			"userId", user.ID)
		return nil, err
	}

	s.logger.InfoKV("User registration completed successfully",
		"username", cmd.Username,
		"userId", user.ID)

	return &SignUpResult{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         user,
	}, nil
}

// Login 用户登录
func (s *AuthCommandService) Login(ctx context.Context, cmd *LoginCmd) (*LoginResult, error) {
	s.logger.DebugKV("Login attempt", "username", cmd.Username)

	conn, err := persistence.NewConn(ctx)
	if err != nil {
		s.logger.ErrorKV("Failed to create database connection for login", "error", err, "username", cmd.Username)
		return nil, err
	}
	defer conn.Close()
	userRepoConn := s.userRepo.WithExecutor(conn)

	// 查找用户
	user, err := userRepoConn.FindByUsername(ctx, cmd.Username)
	if err != nil {
		if errors.Is(err, domainErrors.ErrUserNotFound) {
			s.logger.WarnKV("Login failed - user not found", "username", cmd.Username)
			return nil, errors.New("invalid username or password")
		}
		s.logger.ErrorKV("Failed to find user during login", "error", err, "username", cmd.Username)
		return nil, err
	}

	// 验证密码
	if !auth.CheckPasswordHash(cmd.Password, user.PwdSecret) {
		s.logger.WarnKV("Login failed - invalid password", "username", cmd.Username, "userId", user.ID)
		return nil, errors.New("invalid username or password")
	}

	// 生成token
	accessToken, err := auth.GenerateAccessToken(user.ID, user.Username)
	if err != nil {
		s.logger.ErrorKV("Failed to generate access token", "error", err, "username", cmd.Username, "userId", user.ID)
		return nil, err
	}

	refreshToken, err := auth.GenerateRefreshToken(user.ID, user.Username)
	if err != nil {
		s.logger.ErrorKV("Failed to generate refresh token", "error", err, "username", cmd.Username, "userId", user.ID)
		return nil, err
	}

	s.logger.InfoKV("User logged in successfully", "username", cmd.Username, "userId", user.ID)

	return &LoginResult{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         user,
	}, nil
}

// RefreshToken 刷新token
func (s *AuthCommandService) RefreshToken(ctx context.Context, cmd *RefreshTokenCmd) (*RefreshTokenResult, error) {
	s.logger.Debug("RefreshToken called")

	// 验证刷新token
	claims, err := auth.ValidateToken(cmd.RefreshToken)
	if err != nil {
		s.logger.WarnKV("Invalid refresh token provided", "error", err)
		return nil, errors.New("invalid refresh token")
	}

	// 生成新的访问token
	newAccessToken, err := auth.GenerateAccessToken(claims.UserId, claims.AlternativeID)
	if err != nil {
		s.logger.ErrorKV("Failed to generate new access token",
			"error", err,
			"userId", claims.UserId)
		return nil, err
	}

	s.logger.InfoKV("Access token refreshed successfully", "userId", claims.UserId)

	return &RefreshTokenResult{
		AccessToken: newAccessToken,
	}, nil
}
