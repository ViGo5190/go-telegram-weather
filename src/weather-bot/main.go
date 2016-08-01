package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"os"
	"weather-bot/weather"
)

func main() {
	token := os.Getenv("BOT_TOKEN")

	if token == "" {
		log.Panic("Empty token!")
	}

	weatherToken := os.Getenv("WEATHER_TOKEN")

	if weatherToken == "" {
		log.Panic("Empty weather token!")
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	res, err := bot.RemoveWebhook()

	if err != nil {
		log.Panic(err)
	}
	log.Print(res)

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.Location != nil {
			lat := update.Message.Location.Latitude
			lon := update.Message.Location.Longitude
			resp := weather.GetWeather(weatherToken, lat, lon)

			respTex := "Temp is " + weather.TempToString(resp.Main.Temp) + "C in " + resp.Name

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, respTex)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)

		} else {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			respText := "Please, send me your location"

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, respText)
			msg.ReplyToMessageID = update.Message.MessageID

			bt := tgbotapi.NewKeyboardButtonLocation("Location")

			var btns []tgbotapi.KeyboardButton
			btns = append(btns, bt)

			kbrd := tgbotapi.NewReplyKeyboard(btns);

			msg.ReplyMarkup = kbrd

			bot.Send(msg)
		}

	}
}
