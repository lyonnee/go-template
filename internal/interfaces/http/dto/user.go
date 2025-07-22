package dto

// GetUserResp 获取用户信息响应
type GetUserResp struct {
	User *UserInfo `json:"user"`
}

// UpdateUsernameResp 修改用户名响应
type UpdateUsernameResp struct {
	User *UserInfo `json:"user"`
}
