package model

type UserModel struct {
	SoftDelete_BaseModel
	Username  string `json:"username"`   // Username of the user
	PwdSecret string `json:"pwd_secret"` // Password of the user
	Email     string `json:"email"`      // Email of the user
	Phone     string `json:"phone"`      // Phone number of the user
}
