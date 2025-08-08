package service

type PasswordService interface {
	ComparePassword(secret, password string) (bool, error)
	HashPassword(password string) (string, error)
}

type TokenService interface {
	GenerateToken(userID int64) (string, error)
	ParseToken(token string) (int64, error)
}
