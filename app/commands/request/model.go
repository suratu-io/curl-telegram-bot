package request

import (
    tbot "github.com/go-telegram-bot-api/telegram-bot-api"
    "github.com/suratu-io/curl-telegram-bot/app/commands"
    "net/http"
    "strings"
)

var (
    log = commands.Factory.NewLogger("request")

    // CommandGet is a request command assigned to http "GET" method.
    CommandGet     = NewRequestCommand(http.MethodGet)
    // CommandHead is a request command assigned to http "HEAD" method.
    CommandHead    = NewRequestCommand(http.MethodHead)
    // CommandPost is a request command assigned to http "POST" method.
    CommandPost    = NewRequestCommand(http.MethodPost)
    // CommandPut is a request command assigned to http "PUT" method.
    CommandPut     = NewRequestCommand(http.MethodPut)
    // CommandPatch is a request command assigned to http "PATCH" method.
    CommandPatch   = NewRequestCommand(http.MethodPatch)
    // CommandDelete is a request command assigned to http "DELETE" method.
    CommandDelete  = NewRequestCommand(http.MethodDelete)
    // CommandConnect is a request command assigned to http "CONNECT" method.
    CommandConnect = NewRequestCommand(http.MethodConnect)
    // CommandOptions is a request command assigned to http "OPTIONS" method.
    CommandOptions = NewRequestCommand(http.MethodOptions)
    // CommandTrace is a request command assigned to http "TRACE" method.
    CommandTrace   = NewRequestCommand(http.MethodTrace)
)

// NewRequestCommand returns a new Request Command by the given http method.
func NewRequestCommand(method string) *commands.Command {
    commandName := commands.Name(strings.ToLower(method))
    return commands.NewCommand(commandName, NewRequestCallback(method))
}

// NewRequestCallback returns a new Request Callback by the given http method.
func NewRequestCallback(method string) commands.Callback {
    return func(bot *tbot.BotAPI, msg tbot.Update) error {
        return PerformRequest(method, bot, msg)
    }
}
