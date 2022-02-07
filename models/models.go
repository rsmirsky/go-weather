package models

import (
	"time"
)

// models хранит модельки которые мы будем записывать в базу

type History struct {
	ID               int64
	CreatedAt        time.Time
	TelegramUserID   string
	TelegramUserName string
	Command          string
}
