package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"weather-bot/weather"
)

func removeWebHook(bot *tgbotapi.BotAPI) {
	res, err := (*bot).RemoveWebhook()

	if err != nil {
		log.Panic(err)
	}
	log.Print(res)
}

func handleUpdate(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	if update.Message.Location != nil {
		lat := update.Message.Location.Latitude
		lon := update.Message.Location.Longitude
		resp, err := weather.GetWeather(weatherTkn, lat, lon)
		if (err != nil) {
			log.Panic(err)
		}

		respTex := "Temp is " + weather.TempToString(resp.Main.Temp) + "C in " + resp.Name

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, respTex)
		msg.ReplyToMessageID = update.Message.MessageID

		(*bot).Send(msg)

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

		(*bot).Send(msg)
	}
}

var weatherTkn string
var token string

func main() {
	token = getToken()
	weatherTkn = getWeatherToken()

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)
	removeWebHook(bot)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		go handleUpdate(bot, update)
	}
}
