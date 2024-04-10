package model

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        uint           `gorm:"primarykey"`
	CreatedAt time.Time      `gorm:"column:create_time"`
	UpdatedAt time.Time      `gorm:"column:update_time"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	IsDeleted bool           `gorm:"column:is_deleted"`
}

type User struct {
	BaseModel
	NikeName string `gorm:"type:varchar(20)"`
	Mobile   string `gorm:"index:idx_mobile;unique;type:varchar(11);not null"`
	Password string `gorm:"type:varchar(100);not null"`
	Birthday *time.Time
	Gender   int `gorm:"type:int comment '1-男,0-女'"`
	Role     int `gorm:"default:1;type:int comment '1-普通用户,2-管理员'"`
}
