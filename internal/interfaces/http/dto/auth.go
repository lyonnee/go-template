package dto

// LoginReq 登录请求
type LoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginResp 登录响应
type LoginResp struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// RefreshTokenReq 刷新token请求
type RefreshTokenReq struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

// RefreshTokenResp 刷新token响应
type RefreshTokenResp struct {
	AccessToken string `json:"access_token"`
}

// SignUpReq 注册请求
type SignUpReq struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Password string `json:"password" binding:"required,min=6,max=100"`
	Email    string `json:"email" binding:"required,email"`
	Phone    string `json:"phone" binding:"required"`
}

// SignUpResp 注册响应
type SignUpResp struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	User         *UserInfo `json:"user"`
}

// UserInfo 用户信息
type UserInfo struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

// UpdateUsernameReq 修改用户名请求
type UpdateUsernameReq struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
}
