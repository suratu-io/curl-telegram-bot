package help

import tbot "github.com/go-telegram-bot-api/telegram-bot-api"

func CommandCallback(bot *tbot.BotAPI, msg tbot.Update) error {
    text := "This bot is designed to make possible to send HTTP request from telegram. Available commands: \n" +
        "1. /help\n" +
        "2. /new - start to build new request\n" +
        "3. /curl - make request using standard curl syntax\n"

    res := tbot.NewMessage(msg.Message.Chat.ID, text)

    _, err := bot.Send(res)
    return err
}
