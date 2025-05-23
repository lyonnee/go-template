package model

type BaseModel struct {
	ID        int64 `json:"id"`
	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
}

type SoftDelete_BaseModel struct {
	BaseModel

	IsDeleted bool  `json:"is_deleted"` // Indicates whether the record is soft deleted. If true, the record is considered deleted, and the DeletedAt field stores the timestamp of deletion.
	DeletedAt int64 `json:"deleted_at"`
}
