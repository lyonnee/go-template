package dto

type LoginReq struct {
	PhoneNumber string
	Email       string
	Password    string
}

type LoginResp struct {
	Token string
}

type RefreshTokenReq struct {
	RefreshToken string
}

type RefreshTokenResp struct {
	Token string
}

type SignUpReq struct {
	PhoneNumber string
	Email       string
	Username    string
	Password    string
}

type SignUpResp struct {
}
