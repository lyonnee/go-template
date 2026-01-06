package model

type BaseModel struct {
	ID        uint64 `json:"id" gorm:"primaryKey;autoIncrement"`
	CreatedAt int64  `json:"created_at" gorm:"column:created_at"`
	UpdatedAt int64  `json:"updated_at" gorm:"column:updated_at"`
}

type SoftDelete_BaseModel struct {
	BaseModel

	DeletedAt int64 `json:"deleted_at" gorm:"column:deleted_at;default:0"`
}
