package bot

import (
    tbot "github.com/go-telegram-bot-api/telegram-bot-api"
    "github.com/suratu-io/curl-telegram-bot/app/commands"
    "github.com/suratu-io/curl-telegram-bot/app/commands/help"
)

func (b *Bot) Run(timeout int) error {
    log.Info.Println("update timout:", timeout)

    u := tbot.NewUpdate(0)
    u.Timeout = timeout

    updates, err := b.api.GetUpdatesChan(u)
    if err != nil {
        return err
    }

    for update := range updates {
        b.update(update)
    }

    return nil
}

func (b *Bot) update(update tbot.Update) {
    // Ignore any non-Message Updates.
    if update.Message == nil {
        return
    }

    log.Info.PrintSKV("new message:",
        "update id", update.UpdateID,
        "username", update.Message.From.UserName,
        "text", update.Message.Text,
    )

    if update.Message.IsCommand() {
        name := update.Message.Command()
        commandName := commands.Name(name)
        callback, ok := b.commands.Get(commandName)
        // Sending help message if the command by the given name wasn't found.
        if !ok {
            b.executeCommand(update, help.CommandCallback, help.CommandName)
            return
        }

        log.Info.PrintT("command '{}' found", commandName)
        b.executeCommand(update, callback, commandName)
    }
}

func (b *Bot) executeCommand(update tbot.Update, callback commands.Callback, name commands.Name) {
    err := callback(b.api, update)
    if err != nil {
        log.Error.PrintT("error while executing a command '{}'", name)
        b.sendErrorMessage(update.Message.Chat.ID, err)
        return
    }
}

func (b *Bot) formError(err error) string {
    return "Error trying to execute your command: " + err.Error()
}

func (b *Bot) sendErrorMessage(chatId int64, err error) {
    errMsg := tbot.NewMessage(chatId, b.formError(err))
    _, err = b.api.Send(errMsg)
    if err != nil {
        log.Error.Println("error trying to send an error message:", err)
    }
}
