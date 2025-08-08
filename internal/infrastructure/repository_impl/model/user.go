package model

type UserModel struct {
	SoftDelete_BaseModel

	Username    string `json:"username" db:"username"`           // Username of the user
	PwdSecret   string `json:"pwd_secret" db:"pwd_secret"`       // Password of the user
	Email       string `json:"email" db:"email"`                 // Email of the user
	Phone       string `json:"phone" db:"phone"`                 // Phone number of the user
	LastLoginAt int64  `json:"last_login_at" db:"last_login_at"` // Last login time of the user
}
