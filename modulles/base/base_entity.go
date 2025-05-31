package base

import (
	"time"
)

type BaseEntity struct {
	Id        int64     `json:"id" gorm:"primary_key"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at" autoCreateTime:"true" `
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at" autoUpdateTime:"true" `
}
