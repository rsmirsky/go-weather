package models

import (
	"time"
)

type History struct {
	ID               int64
	CreatedAt        time.Time
	TelegramUserID   string
	TelegramUserName string
	Command          string
}
