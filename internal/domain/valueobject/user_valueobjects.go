package valueobject

import (
	"errors"
	"regexp"
	"strings"
)

// Email 邮箱值对象
type Email struct {
	value string
}

// NewEmail 创建新的邮箱值对象
func NewEmail(email string) (*Email, error) {
	email = strings.TrimSpace(email)
	if email == "" {
		return nil, errors.New("email cannot be empty")
	}

	// 简单的邮箱格式验证
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return nil, errors.New("invalid email format")
	}

	return &Email{value: email}, nil
}

// Value 获取邮箱值
func (e Email) Value() string {
	return e.value
}

// String 实现 Stringer 接口
func (e Email) String() string {
	return e.value
}

// Phone 手机号值对象
type Phone struct {
	value string
}

// NewPhone 创建新的手机号值对象
func NewPhone(phone string) (*Phone, error) {
	phone = strings.TrimSpace(phone)
	if phone == "" {
		return nil, errors.New("phone cannot be empty")
	}

	// 简单的手机号格式验证 (中国手机号)
	phoneRegex := regexp.MustCompile(`^1[3-9]\d{9}$`)
	if !phoneRegex.MatchString(phone) {
		return nil, errors.New("invalid phone format")
	}

	return &Phone{value: phone}, nil
}

// Value 获取手机号值
func (p Phone) Value() string {
	return p.value
}

// String 实现 Stringer 接口
func (p Phone) String() string {
	return p.value
}

// Username 用户名值对象
type Username struct {
	value string
}

// NewUsername 创建新的用户名值对象
func NewUsername(username string) (*Username, error) {
	username = strings.TrimSpace(username)
	if username == "" {
		return nil, errors.New("username cannot be empty")
	}

	if len(username) < 3 || len(username) > 50 {
		return nil, errors.New("username length must be between 3 and 50 characters")
	}

	// 用户名只能包含字母、数字和下划线
	usernameRegex := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
	if !usernameRegex.MatchString(username) {
		return nil, errors.New("username can only contain letters, numbers and underscores")
	}

	return &Username{value: username}, nil
}

// Value 获取用户名值
func (u Username) Value() string {
	return u.value
}

// String 实现 Stringer 接口
func (u Username) String() string {
	return u.value
}
