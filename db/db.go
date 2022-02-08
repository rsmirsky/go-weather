package db

import (
	"weather/models"
	req "weather/models/requests"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// gorm.DB используем только в этом пакете

var (
	db *gorm.DB
)

func Connect(dsn string) error {
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	db = d
	return nil
}

func RunMigrates() {
	db.AutoMigrate(&models.History{})
}

func CreateHistory(r req.NewHistory) error {
	history := models.History{
		TelegramChatID:   r.TelegramChatID,
		TelegramUserName: r.TelegramUserName,
		Command:          r.Command,
		IsCity:           r.IsCity,
	}

	return db.Create(&history).Error
}
