package entity

type User struct {
	ID        int64
	CreatedAt int64
	UpdatedAt int64

	Username string
	Password string
	Email    string
	Phone    string

	IsDeleted bool
	DeletedAt int64
}
