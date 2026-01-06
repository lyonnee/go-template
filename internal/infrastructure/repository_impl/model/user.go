package model

type UserModel struct {
	SoftDelete_BaseModel

	Username    string `json:"username" gorm:"column:username"`
	PwdSecret   string `json:"pwd_secret" gorm:"column:pwd_secret"`
	Email       string `json:"email" gorm:"column:email"`
	Phone       string `json:"phone" gorm:"column:phone"`
	LastLoginAt int64  `json:"last_login_at" gorm:"column:last_login_at"`
}

func (UserModel) TableName() string {
	return "users"
}
