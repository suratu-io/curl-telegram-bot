package bot

import (
    tbot "github.com/go-telegram-bot-api/telegram-bot-api"
    "github.com/suratu-io/curl-telegram-bot/app/commands"
    "github.com/suratu-io/curl-telegram-bot/app/logger"
)

type (
    // Bot is a wrapper for tbot.BotAPI that stricts and simplifies
    // its functionality.
    Bot struct {
        RunBotRepo
        api      *tbot.BotAPI
        commands commands.CallbackMap
    }
)

var (
    log = logger.Factory.NewLogger("bot")
)

// NewBot initializes bot api and returns a new *Bot.
func NewBot(token string, commands commands.CallbackMap) (*Bot, error) {
    api, err := tbot.NewBotAPI(token)
    if err != nil {
        return nil, err
    }

    return &Bot{api: api, commands: commands}, nil
}
