package request

import (
	tbot "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/suratu-io/curl-telegram-bot/app/commands"
	"net/http"
	"strings"
)

var (
	CommandGet 		= NewRequestCommand(http.MethodGet)
	CommandHead 	= NewRequestCommand(http.MethodHead)
	CommandPost 	= NewRequestCommand(http.MethodPost)
	CommandPut 		= NewRequestCommand(http.MethodPut)
	CommandPatch 	= NewRequestCommand(http.MethodPatch)
	CommandDelete 	= NewRequestCommand(http.MethodDelete)
	CommandConnect 	= NewRequestCommand(http.MethodConnect)
	CommandOptions 	= NewRequestCommand(http.MethodOptions)
	CommandTrace 	= NewRequestCommand(http.MethodTrace)
)

func NewRequestCommand(method string) *commands.Command {
	commandName := commands.Name(strings.ToLower(method))
	return commands.NewCommand(commandName, NewRequestCallback(method))
}

func NewRequestCallback(method string) commands.Callback {
	return func(bot *tbot.BotAPI, msg tbot.Update) error {
		return PerformRequest(method, bot, msg)
	}
}
