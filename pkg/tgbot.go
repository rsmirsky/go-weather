package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"weather/models"

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

			// TODO: принимать город в формате kyiv/Kyiv 
			// регистронезависимо
			// когда сравниваешь строки - обычно это делается в LowerCase

			cityUser, err := cityList.GetCityId(update.Message.Text)
			if err != nil {
				// wrong city
				// send message to telegram "unknown error"
				
				
				// TODO вынести в константу
				
				
				// TODO: Правильно назвать переменные

				// TODO: Понять почему Одесса в ответе возвращается с одной буквой (или в запросе с двумя)
				
				// TODO: Поудалять ненужные комментарии которые не относятся к коду

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, wrongCityMessage)
				bot.Send(msg)
			} else {
				//fmt.Println(cityUser)

		
				getWeatherStruct := GetWeather(cityUser)
         		
			
				// TODO: Возвращать больше информации о погоде, насколько возможно
				
				message := fmt.Sprintf(weatherMessage, getWeatherStruct.Name, getWeatherStruct.GetCelsius(), getWeatherStruct.GetCelsiusMin(),getWeatherStruct.GetCelsiusMax(), getWeatherStruct.GetClouds())
			 	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
			 	bot.Send(msg)
				
			}
            




		// msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		// // Extract the command from the Message.
		// switch update.Message.Command() {
		// case "help":
		// 	msg.Text = "I understand /sayhi and /status."
		// case "start":
		// 	msg.Text = "Hi :)"
		// case "status":
		// 	msg.Text = "I'm ok."
		// default:
		// 	msg.Text = "I don't know that command"
		// }

		// if _, err := bot.Send(msg); err != nil {
		// 	log.Panic(err)
		// }

			

			

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


// TODO: Возвращать структуру Weather
// TODO: возвращать ошибку вторым аргументом
func GetWeather(cityUser float64) (weather *models.MainWeather) {
	
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

	weather = &models.MainWeather{}
	err = json.Unmarshal([]byte(b), weather)
	if err != nil {
		log.Fatalln(err, "unmarshal error")
	}

	return weather
	

	//fmt.Println("Текущая температура в Киеве", tempC)

	//var tokenBot string = "5294861035:AAFBhY-RL88hOkAccUW_Xz7Jh83a6Iwpa8Y"

	//ConnectToBot(tokenBot, tempC)

}




const( 
	wrongCityMessage = "Enter the name of your city in Latin letters.\nKyiv, Dnipro, Kharkiv, Odessa..."
	weatherMessage = `In the city %s 
	Temperature: %g°C
	Min temperature: %g°C
	Max temperature: %g°C
	Clouds: %s
	`
)
