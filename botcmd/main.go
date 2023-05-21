/*
SIMPLE SLACK BOT THAT RESPONDS TO 2 CURRENT QUERIES: GREETINGS AND AGE
*/

package botcmd

import (
	"fmt"
	"strconv"

	"github.com/moh3a/slack-go-bots/shared"
	"github.com/shomali11/slacker"
)

func Run() {
	bot := shared.NewSlackerClient()

	botCommand(
		bot,
		Command{
			formatted_prompt: "My year of birth is <year>",
			description:      "Year Of Birth Calculator",
			examples:         []string{"My year of birth is 2020"},
			handler: func(bc slacker.BotContext, r slacker.Request, w slacker.ResponseWriter) {
				year := r.Param("year")
				yob, err := strconv.Atoi(year)
				if err != nil {
					fmt.Println("Error!")
				}
				age := 2023 - yob
				reply := fmt.Sprintf("Your age is %d!", age)
				w.Reply(reply)
			},
		},
	)

	botCommand(
		bot,
		Command{
			formatted_prompt: "My name is <name>",
			description:      "Greetings From Our Bot!",
			examples:         []string{"My name is Alex"},
			handler: func(bc slacker.BotContext, r slacker.Request, w slacker.ResponseWriter) {
				name := r.Param("name")
				reply := fmt.Sprintf("Hello %s. My name is bot. reglini.bot.", name)
				w.Reply(reply)
			},
		},
	)

	botCommand(
		bot,
		Command{
			formatted_prompt: "? <message>",
			description:      "Send any question to wolfram",
			examples:         []string{"? Who is the president of the world?"},
			handler: func(bc slacker.BotContext, r slacker.Request, w slacker.ResponseWriter) {
				query := r.Param("message")
				fmt.Printf(query) // todo
			},
		},
	)
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
