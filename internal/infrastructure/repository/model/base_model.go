package model

type BaseModel struct {
	ID        int64 `json:"id"`
	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
}

type SoftDelete_BaseModel struct {
	BaseModel

	DeletedAt int64 `json:"deleted_at"`
}
