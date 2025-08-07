package entity

import (
	"errors"
	"regexp"
	"strings"
	"time"

	"github.com/lyonnee/go-template/internal/infrastructure/auth"
	"github.com/lyonnee/go-template/pkg/di"
)

type User struct {
	ID        int64
	CreatedAt int64
	UpdatedAt int64

	Username  string
	PwdSecret string
	Email     string
	Phone     string

	DeletedAt int64
}

func NewUser(username, pwd, email, phone string) (*User, error) {
	now := time.Now().Unix()

	if err := validateUsername(username); err != nil {
		return nil, err
	}
	if err := validateEmail(email); err != nil {
		return nil, err
	}
	if err := validatePhone(phone); err != nil {
		return nil, err
	}
	if err := validatePassword(pwd); err != nil {
		return nil, err
	}

	pwdSecret, err := auth.HashPassword(pwd)
	if err != nil {
		return nil, err
	}

	u := &User{
		CreatedAt: now,
		UpdatedAt: now,

		DeletedAt: 0,

		Username:  username,
		PwdSecret: pwdSecret,
		Email:     email,
		Phone:     phone,
	}

	return u, nil
}

func (u *User) Login(pwd string) (string, string, error) {
	// validate password
	if !auth.CheckPasswordHash(pwd, u.PwdSecret) {
		return "", "", errors.New("invalid username or password")
	}

	return u.BuildToken()
}

func (u *User) BuildToken() (string, string, error) {
	// build access token and refresh token
	jwtManager := di.Get[*auth.JWTManager]()

	accessToken, err := jwtManager.GenerateAccessToken(u.ID, u.Username)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := jwtManager.GenerateRefreshToken(u.ID, u.Username)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func validateUsername(username string) error {
	username = strings.TrimSpace(username)
	if username == "" {
		return errors.New("username cannot be empty")
	}

	if len(username) < 3 || len(username) > 50 {
		return errors.New("username length must be between 3 and 50 characters")
	}

	// 用户名只能包含字母、数字和下划线
	usernameRegex := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
	if !usernameRegex.MatchString(username) {
		return errors.New("username can only contain letters, numbers and underscores")
	}

	return nil
}

func validateEmail(email string) error {
	email = strings.TrimSpace(email)
	if email == "" {
		return errors.New("email cannot be empty")
	}

	// 简单的邮箱格式验证
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return errors.New("invalid email format")
	}

	return nil
}

func validatePhone(phone string) error {
	phone = strings.TrimSpace(phone)
	if phone == "" {
		return errors.New("phone cannot be empty")
	}

	// 简单的手机号格式验证 (中国手机号)
	phoneRegex := regexp.MustCompile(`^1[3-9]\d{9}$`)
	if !phoneRegex.MatchString(phone) {
		return errors.New("invalid phone format")
	}

	return nil
}

func validatePassword(pwd string) error {
	pwd = strings.TrimSpace(pwd)
	if pwd == "" {
		return errors.New("password cannot be empty")
	}
	if len(pwd) < 6 || len(pwd) > 50 {
		return errors.New("password length must be between 6 and 50 characters")
	}
	return nil
}
