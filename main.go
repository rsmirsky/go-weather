package main

import (
	"fmt"
	"log"
	req "weather/models/requests"
	"weather/openweathermap"

	"weather/db"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("error read in config: ", err)
		return
	}

	if err := db.Connect(viper.GetString("DSN_MYSQL")); err != nil {
		fmt.Println("error connect to db:", err)
		return
	}

	db.RunMigrates()

	openweathermap.Init(viper.GetString("OPENWEATHERMAP_APPID"))

	RunBot(viper.GetString("TELEGRAM_BOT_TOKEN"))
}

func RunBot(botToken string) {

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	cityList, err := openweathermap.GetCityList()
	if err != nil {
		fmt.Println("Failed to get city list: ", err)
		return
	}
    var IsCityVal bool = true
	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			cityId, err := cityList.GetCityId(update.Message.Text)
			if err != nil {

				// Wrong city
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, wrongCityMessage)
				bot.Send(msg)
				IsCityVal = false
			} else {

				// Correct city
				weather, err := openweathermap.GetWeather(cityId)
				if err != nil {
					log.Println("get weather error: ", err)
					continue
				}

				message := fmt.Sprintf(weatherMessage, update.Message.Text, weather.GetCelsius(), weather.GetCelsiusMin(), weather.GetCelsiusMax(), weather.GetClouds())
				IsCityVal = true
				//models.SaveWeatherQueryToDB(update.Message.Text string)
				//userHistory := &models.History{}
				//userHistory := &models.History{}
				//userHistory.TelegramUserID =
				//models.SaveWeatherQueryToDB()

				

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
				bot.Send(msg)

			}
		}
		history := req.NewHistory{
			TelegramChatID:   update.Message.Chat.ID,
			TelegramUserName: update.Message.From.UserName,
			Command:          update.Message.Text,
			IsCity:           IsCityVal,
		}
		db.CreateHistory(history)
	}
}

const (
	wrongCityMessage = "Enter the name of your city in Latin letters.\nKyiv, Dnipro, Kharkiv, Odessa..."
	weatherMessage   = `In the city %s 
	Temperature: %g°C
	Min temperature: %g°C
	Max temperature: %g°C
	Clouds: %s
	`
)
