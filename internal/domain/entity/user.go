package entity

import (
	"regexp"
	"strings"
	"time"

	"github.com/lyonnee/go-template/internal/domain/errors"
	"github.com/lyonnee/go-template/pkg/util"
)

type User struct {
	ID        uint64
	CreatedAt int64
	UpdatedAt int64
	DeletedAt int64

	Username    string
	PwdSecret   string
	Email       string
	Phone       string
	LastLoginAt int64
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

	pwdSecret, err := util.HashPassword(pwd)
	if err != nil {
		return nil, err
	}

	u := &User{
		CreatedAt:   now,
		UpdatedAt:   now,
		LastLoginAt: 0,
		DeletedAt:   0,

		Username:  username,
		PwdSecret: pwdSecret,
		Email:     email,
		Phone:     phone,
	}

	return u, nil
}

func (u *User) Login(pwd string) error {
	if u.DeletedAt > 0 {
		return errors.ErrUserDeleted
	}

	// validate password
	if err := util.ComparePassword(pwd, u.PwdSecret); err != nil {
		return errors.ErrInvalidPassword
	}

	// update last login time
	u.LastLoginAt = time.Now().Unix()

	return nil
}

func (u *User) UpdatePassword(pwd string) error {
	if err := validatePassword(pwd); err != nil {
		return err
	}

	pwdSecret, err := util.HashPassword(pwd)
	if err != nil {
		return err
	}

	u.PwdSecret = pwdSecret
	u.UpdatedAt = time.Now().Unix()

	return nil
}

func validateUsername(username string) error {
	username = strings.TrimSpace(username)
	if username == "" {
		return errors.ErrInvalidUsername
	}

	if len(username) < 3 || len(username) > 50 {
		return errors.ErrInvalidUsername
	}

	// 用户名只能包含字母、数字和下划线
	usernameRegex := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
	if !usernameRegex.MatchString(username) {
		return errors.ErrInvalidUsernameFormat
	}

	return nil
}

func validateEmail(email string) error {
	email = strings.TrimSpace(email)
	if email == "" {
		return errors.ErrInvalidEmail
	}

	// 简单的邮箱格式验证
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return errors.ErrInvalidEmailFormat
	}

	return nil
}

func validatePhone(phone string) error {
	phone = strings.TrimSpace(phone)
	if phone == "" {
		return errors.ErrInvalidPhone
	}

	// 简单的手机号格式验证 (中国手机号)
	phoneRegex := regexp.MustCompile(`^1[3-9]\d{9}$`)
	if !phoneRegex.MatchString(phone) {
		return errors.ErrInvalidPhoneFormat
	}

	return nil
}

func validatePassword(pwd string) error {
	pwd = strings.TrimSpace(pwd)
	if pwd == "" {
		return errors.ErrInvalidPassword
	}
	if len(pwd) < 6 || len(pwd) > 50 {
		return errors.ErrInvalidPassword
	}
	return nil
}
