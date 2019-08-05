package help

import "github.com/suratu-io/curl-telegram-bot/app/commands"

const (
    CommandName = "help"
)

var (
    Command = commands.NewCommand(CommandName, CommandCallback)
)
