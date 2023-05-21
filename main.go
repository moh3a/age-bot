package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/moh3a/slack-go-bots/botcmd"
	"github.com/moh3a/slack-go-bots/upload"
)

func main() {
	enverr := godotenv.Load(".env")
	if enverr != nil {
		log.Fatalf("Error loading environment variables file")
	}

	fileArr := []string{"./assets/aliexpress-icon.png", "./icon-512x512.png"}

	upload.Upload(fileArr) // SIMPLE SLACK FILE UPLOAD BOT
	botcmd.Run()           // SIMPLE SLACK BOT THAT RESPONDS TO 2 CURRENT QUERIES: GREETINGS AND AGE
}
