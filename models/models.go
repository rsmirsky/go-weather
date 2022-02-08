package models

import (
	//"fmt"
	"time"
)

// models хранит модельки которые мы будем записывать в базу

type History struct {
	ID               int64 `gorm:"primaryKey"`
	CreatedAt        time.Time
	TelegramChatID   int64
	TelegramUserName string
	Command          string
	IsCity           bool
}
