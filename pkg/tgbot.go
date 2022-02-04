package pkg

import (
	"fmt"
	"log"
	"weather/models"
	"encoding/json"
	"io/ioutil"
	//"log"
	"net/http"
	//"weather/models"


	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func ConnectToBot() {

    var tokenBot string = "5294861035:AAFBhY-RL88hOkAccUW_Xz7Jh83a6Iwpa8Y"

	bot, err := tgbotapi.NewBotAPI(tokenBot)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	cityList, err := models.GetCityList()
	if err != nil {
		fmt.Println("Failed to get city list: ", err)
		return
	}

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			// TODO: Добавить обработку команды /start

			cityUser, err := cityList.GetCityId(update.Message.Text)
			if err != nil {
				// wrong city
				// send message to telegram "unknown error"
				msgTemp := "The city is incorrect.\nThe name of the city is written in Latin letters.\n'Kyiv, Dnipro, Kharkiv, Odessa..."
				
				// TODO: исправить Название города Kyiv
				// TODO: Понять почему Одесса в ответе возвращается с одной буквой (или в запросе с двумя)
				
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgTemp)
				bot.Send(msg)
			} else {
				//fmt.Println(cityUser)
				userCityTemp, userCity := ConnectToApi(cityUser)
             	msgTemp := fmt.Sprintf("in the city of %s %g degrees C",userCity, userCityTemp)
			 	msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgTemp)
			 	bot.Send(msg)
				
			}


			

			

			// switch update.Message.Text {

			// case "temp":

			// 	cityNew := cityList.GetCityId(userCity)
			// 	msgTemp := fmt.Sprintf("Сейчас в Киеве %g ", c)
			// 	msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgTemp)
			// 	bot.Send(msg)

			// default:

			// 	msgDefault := fmt.Sprintf("Сейчас в Киеве %g ", c)
			// 	msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgDefault)
			// 	bot.Send(msg)
			//}
		}

		//msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		//msg.ReplyToMessageID = update.Message.MessageID

		//bot.Send(msg)
	}

}

func ConnectToApi(cityUser float64) (float64, string) {
	
    var cityNum float64 = cityUser

	var link = fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?id=%.0f&appid=90002db5f9ad4366f2430e6ec2abf999", cityNum)
	fmt.Println("http.Get url: ", link)
	resp, err := http.Get(link)
	if err != nil {
		log.Fatalln(err, "http.Get error")
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err, "readall error")
	} 

	// fmt.Println("=========")
	// fmt.Println(string(b))
	// fmt.Println("=========")

	weather := &models.MainWeather{}
	err = json.Unmarshal([]byte(b), weather)
	if err != nil {
		log.Fatalln(err, "unmarshal error")
	}

	tempC, userCity := weather.GetC()
	return tempC, userCity

	//fmt.Println("Текущая температура в Киеве", tempC)

	//var tokenBot string = "5294861035:AAFBhY-RL88hOkAccUW_Xz7Jh83a6Iwpa8Y"

	//ConnectToBot(tokenBot, tempC)


}
