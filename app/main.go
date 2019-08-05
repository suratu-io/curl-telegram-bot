package main

import (
    "github.com/suratu-io/curl-telegram-bot/app/bot"
    "github.com/suratu-io/curl-telegram-bot/app/commands"
    "github.com/suratu-io/curl-telegram-bot/app/commands/help"
    "github.com/suratu-io/curl-telegram-bot/app/commands/request"
    "github.com/suratu-io/curl-telegram-bot/app/config"
    "github.com/suratu-io/curl-telegram-bot/app/logger"
)

var (
    log = logger.Factory.NewLogger("main")
    // Mapping commands into a map to make command selection easier.
    cmds = commands.NewCallbackMap(
        help.Command,
        request.CommandGet,
        request.CommandHead,
        request.CommandPost,
        request.CommandPut,
        request.CommandPatch,
        request.CommandDelete,
        request.CommandConnect,
        request.CommandOptions,
        request.CommandTrace,
    )
)

func main() {
    // Creating and setting up a new bot api client.
    b, err := bot.NewBot(config.BotApiToken, config.ProductionMode, cmds)
    if err != nil {
        log.Panic.Println("error trying to initialize a new bot")
        panic(err)
    }

    // Run is going to loop a continues chan that will block
    // the further execution of main func.
    err = b.Run(config.UpdateTimeout)
    if err != nil {
        log.Panic.Println("error trying to run bot")
        panic(err)
    }
}
