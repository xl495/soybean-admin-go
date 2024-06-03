package model

import (
	"gorm.io/gorm"
	"time"
)

type GormBase struct {
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time      `gorm:"column:createdTime;" json:"createdTime"`
	UpdatedAt time.Time      `gorm:"column:updatedTime;" json:"updatedTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deletedTime;index" json:"deletedTime,omitempty"`
}
