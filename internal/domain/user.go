package domain

type UserRepository interface {
	FindById(userId int64) (*User, error)
}

type User struct {
	UserID   int64
	Username string
}
