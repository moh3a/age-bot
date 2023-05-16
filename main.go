package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/moh3a/slack-go-bots/bot"
	"github.com/moh3a/slack-go-bots/upload"
)

func main() {
	enverr := godotenv.Load(".env")
	if enverr != nil {
		log.Fatalf("Error loading environment variables file")
	}

	fileArr := []string{"./assets/aliexpress-icon.png", "./icon-512x512.png"}

	upload.Upload(fileArr) // SIMPLE SLACK FILE UPLOAD BOT
	bot.Run()              // SIMPLE SLACK BOT THAT RESPONDS TO 2 CURRENT QUERIES: GREETINGS AND AGE
}
