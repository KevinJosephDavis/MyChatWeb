package model

import (
	"time"
)

type Friend struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"not null"`
	FriendID  uint   `gorm:"not null"`
	Remark    string `gorm:"size:50"`
	Status    int8   `gorm:"default:1"`
	CreatedAt time.Time

	User   User `gorm:"foreignKey:UserID"`
	Friend User `gorm:"foreignKey:FriendID"`
}
