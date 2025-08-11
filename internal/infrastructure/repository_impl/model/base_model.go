package model

type BaseModel struct {
	ID        uint64 `json:"id" db:"id"` // Primary key ID
	CreatedAt int64  `json:"created_at" db:"created_at"`
	UpdatedAt int64  `json:"updated_at" db:"updated_at"`
}

type SoftDelete_BaseModel struct {
	BaseModel

	DeletedAt int64 `json:"deleted_at" db:"deleted_at"` // Soft delete timestamp
}
