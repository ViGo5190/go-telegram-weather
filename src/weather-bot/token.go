package main

import (
	"os"
	"log"
)

func getToken() string {
	token := os.Getenv("BOT_TOKEN")

	if token == "" {
		log.Panic("Empty token!")
	}
	return token
}

func getWeatherToken() string {
	weatherToken := os.Getenv("WEATHER_TOKEN")

	if weatherToken == "" {
		log.Panic("Empty weather token!")
	}
	return weatherToken
}