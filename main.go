package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
)

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

// remove tokens before pushing to gh
var SLACK_APP_TOKEN = "{{SLACK_APP_TOKEN}}"
var SLACK_BOT_TOKEN = "{{SLACK_BOT_TOKEN}}"

func main() {
	os.Setenv("SLACK_BOT_TOKEN", SLACK_BOT_TOKEN)
	os.Setenv("SLACK_APP_TOKEN", SLACK_APP_TOKEN)

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("my year of birth is <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
		Examples:    []string{"my year of birth is 2020"},
		Handler: func(bc slacker.BotContext, r slacker.Request, w slacker.ResponseWriter) {
			year := r.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				fmt.Println("Error!")
			}
			age := 2023 - yob
			reply := fmt.Sprintf("age is %d", age)
			w.Reply(reply)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
