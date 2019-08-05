package help

import tbot "github.com/go-telegram-bot-api/telegram-bot-api"

func CommandCallback(bot *tbot.BotAPI, msg tbot.Update) error {
    text := "This bot is designed to make possible to send HTTP request from telegram. Available commands: \n" +
        "1. /help - show this message\n" +
        "2. /{http_method} url - start to build new request\n"

    res := tbot.NewMessage(msg.Message.Chat.ID, text)

    _, err := bot.Send(res)
    return err
}
