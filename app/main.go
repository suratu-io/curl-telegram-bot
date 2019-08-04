package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/olehan/kek"
	"os"
)

var (
	log = kek.NewLogger("main")
)

const (
	help   = "/help"
	get = "/get"
	post = "/post"
	put = "/put"
	delete = "/delete"
	update = "/update"
	head = "/head"
	trace = "/trace"
	patch = "/patch"
	connect = "/connect"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("tgbot_token"))
	if err != nil {
		log.Panic.Println("error trying to connect to the the telegram bot api")
		panic(err)
	}

	bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Succ.PrintT("{} {}", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

		switch msg.Text {
		case help:
			sendHelpMessage(msg, bot)
		case newReq:
			buildNewRequest(msg, bot)
		}
	}
}

func sendHelpMessage(msg tgbotapi.MessageConfig, bot *tgbotapi.BotAPI) {
	text := "This bot is designed to make possible to send HTTP request from telegram. Available commands: \n" +
		"1. /help\n" +
		"2. /new - start to build new request\n" +
		"3. /curl - make request using standard curl syntax\n"
	response := tgbotapi.NewMessage(msg.ChatID, text)

    if _, err := bot.Send(response); err != nil {
        log.Error.Println("error during help send")
    }
}

func buildNewRequest(msg tgbotapi.MessageConfig, bot *tgbotapi.BotAPI) {
	text := "Enter request url:"
	textResponse := tgbotapi.NewMessage(msg.ChatID, text)


    if _, err := bot.Send(textResponse); err != nil {
        log.Error.PrintT("error during new request forming: {}", err)
    }
}
