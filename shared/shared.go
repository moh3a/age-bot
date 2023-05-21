package shared

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/shomali11/slacker"
)

func NewSlackerClient() *slacker.Slacker {
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))
	go printCommandEvents(bot.CommandEvents())

	botctx, botcancel := context.WithCancel(context.Background())
	defer botcancel()

	boterr := bot.Listen(botctx)
	if boterr != nil {
		log.Fatal(boterr)
	}

	return bot
}

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}
