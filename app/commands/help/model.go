package help

import "github.com/suratu-io/curl-telegram-bot/app/commands"

const (
    // CommandName is the name for command "help".
    CommandName = "help"
)

var (
    // Command is the composed Command for "help".
    // "help" sends a help message that has usage instructions
    // of the curl bot.
    Command = commands.NewCommand(CommandName, CommandCallback)
)
