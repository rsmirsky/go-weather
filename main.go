package main

import (
	"fmt"
	"log"
	"weather/openweathermap"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/spf13/viper"
)

func main() {

	// db url: root:@tcp(127.0.0.1:3306)/weather?charset=utf8mb4&parseTime=True&loc=Local

	viper.SetConfigFile(".env")
	viper.ReadInConfig()

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

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			cityId, err := cityList.GetCityId(update.Message.Text)
			if err != nil {

				// Wrong city
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, wrongCityMessage)
				bot.Send(msg)
			} else {

				// Correct city
				weather, err := openweathermap.GetWeather(cityId)
				if err != nil {
					log.Println("get weather error: ", err)
					continue
				}

				message := fmt.Sprintf(weatherMessage, update.Message.Text, weather.GetCelsius(), weather.GetCelsiusMin(), weather.GetCelsiusMax(), weather.GetClouds())
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
				bot.Send(msg)
			}
		}
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
