/*
SIMPLE SLACK BOT THAT RESPONDS TO 2 CURRENT QUERIES: GREETINGS AND AGE
*/

package slacker

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
)

func RunBot() {
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))
	go printCommandEvents(bot.CommandEvents())

	botCommand(
		bot,
		Command{
			formatted_prompt: "my year of birth is <year>",
			description:      "yob calculator",
			examples:         []string{"my year of birth is 2020"},
			handler: func(bc slacker.BotContext, r slacker.Request, w slacker.ResponseWriter) {
				year := r.Param("year")
				yob, err := strconv.Atoi(year)
				if err != nil {
					fmt.Println("Error!")
				}
				age := 2023 - yob
				reply := fmt.Sprintf("age is %d", age)
				w.Reply(reply)
			},
		},
	)

	botCommand(
		bot,
		Command{
			formatted_prompt: "hello, my name is <name>",
			description:      "greetings from out bot!",
			examples:         []string{"hello, my name is mohamed"},
			handler: func(bc slacker.BotContext, r slacker.Request, w slacker.ResponseWriter) {
				name := r.Param("name")
				reply := fmt.Sprintf("hello %s. My name is bot. reglini.bot.", name)
				w.Reply(reply)
			},
		},
	)

	botctx, botcancel := context.WithCancel(context.Background())
	defer botcancel()

	boterr := bot.Listen(botctx)
	if boterr != nil {
		log.Fatal(boterr)
	}
}

type Command struct {
	formatted_prompt string
	description      string
	examples         []string
	handler          func(slacker.BotContext, slacker.Request, slacker.ResponseWriter)
}

func botCommand(bot *slacker.Slacker, command Command) {
	bot.Command(command.formatted_prompt, &slacker.CommandDefinition{
		Description: command.description,
		Examples:    command.examples,
		Handler:     command.handler,
	})
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
