package model

import (
	"time"
)

type User struct {
	ID          int64      `gorm:"primaryKey;autoIncrement" json:"id"`
	AccountID   int64      `gorm:"uniqueIndex;not null" json:"account_id"` // 对外展示ID
	Username    string     `gorm:"uniqueIndex;size:50;not null" json:"username"`
	Password    string     `gorm:"size:255;not null" json:"-"` // 密码字段不导出为JSON
	Email       string     `gorm:"size:100;uniqueIndex" json:"email"`
	Nickname    string     `gorm:"size:50" json:"nickname"`
	Avatar      string     `gorm:"size:255" json:"avatar"`
	Status      int8       `gorm:"default:1" json:"status"` // 1: active, 0: inactive
	LastLoginAt *time.Time `gorm:"column:last_login_at" json:"last_login_at,omitempty"`
	CreatedAt   time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
}
