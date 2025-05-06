package dto

type LoginReq struct {
	PhoneNumber string
	Email       string
	Password    string
}

type SignUpReq struct {
	PhoneNumber string
	Email       string
	Username    string
	Password    string
}
