package service

type TokenService interface {
	GenerateToken(userID int64) (string, error)
	ParseToken(token string) (int64, error)
}
