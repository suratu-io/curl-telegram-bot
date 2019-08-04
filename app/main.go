package main

import (
    "github.com/go-telegram-bot-api/telegram-bot-api"
    "github.com/olehan/kek"
    "os"
)

var (
    log = kek.NewLogger("main")
)

func main() {
    bot, err := tgbotapi.NewBotAPI(os.Getenv("tgbot_token"))
    if err != nil {
        log.Panic.Println("error trying to connect to the the telegram bot api")
        panic(err)
    }

    bot.Debug = true
    log.Succ.Println("app successfully connected to the telegram bot api")
}
