package application

type LoginCmd struct {
	PhoneNumber string
	Email       string
	Password    string
}

type AuthService interface {
	Login(cmd *LoginCmd) error
}
