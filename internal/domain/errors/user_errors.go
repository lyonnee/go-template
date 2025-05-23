package errors

import (
	"errors"
)

// 用户相关错误
var (
	ErrUserNotFound     = errors.New("user not found")
	ErrUsernameTaken    = errors.New("username already taken")
	ErrEmailTaken       = errors.New("email already taken")
	ErrPhoneTaken       = errors.New("phone number already taken")
	ErrInvalidUserInput = errors.New("invalid user input")
)
