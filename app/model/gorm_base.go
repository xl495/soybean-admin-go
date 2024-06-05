package model

import (
	"gorm.io/gorm"
	"time"
)

type GormBase struct {
	ID        uint           `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt time.Time      `gorm:"column:createdTime;" json:"createTime"`
	UpdatedAt time.Time      `gorm:"column:updatedTime;" json:"updateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deletedTime;index" json:"deleteTime,omitempty"`
}
