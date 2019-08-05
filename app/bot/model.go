package bot

import (
    tbot "github.com/go-telegram-bot-api/telegram-bot-api"
    "github.com/suratu-io/curl-telegram-bot/app/commands"
    "github.com/suratu-io/curl-telegram-bot/app/logger"
)

type (
    Bot struct {
        RunBotRepo
        api      *tbot.BotAPI
        commands commands.CallbackMap
    }
)

var (
    log = logger.Factory.NewLogger("bot")
)

func NewBot(token string, debug bool, commands commands.CallbackMap) (*Bot, error) {
    api, err := tbot.NewBotAPI(token)
    if err != nil {
        return nil, err
    }

    api.Debug = debug

    return &Bot{api: api, commands: commands}, nil
}
