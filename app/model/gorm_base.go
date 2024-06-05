package model

import (
	"gorm.io/gorm"
	"time"
)

type GormBase struct {
	ID        uint           `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt time.Time      `gorm:"column:created_time;comment: 创建时间" json:"createTime"`
	UpdatedAt time.Time      `gorm:"column:updated_time;comment: 更新时间" json:"updateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_time;index" json:"deleteTime,omitempty"`
}
