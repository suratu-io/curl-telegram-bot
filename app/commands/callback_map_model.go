package commands

import (
    tbot "github.com/go-telegram-bot-api/telegram-bot-api"
)

type (
    Name string
    Callback func(bot *tbot.BotAPI, msg tbot.Update) error
    CallbackMap map[Name]Callback
)

func NewCallbackMap(commands ...*Command) CallbackMap {
    commandMap := CallbackMap{}
    for _, command := range commands {
        commandMap.SetFromCommand(command)
    }
    return commandMap
}

func (c CallbackMap) Get(name Name) (callback Callback, ok bool) {
    callback, ok = c[name]
    return
}

func (c CallbackMap) Set(name Name, callback Callback) CallbackMap {
    c[name] = callback
    return c
}

func (c CallbackMap) SetFromCommand(cmd *Command) CallbackMap {
    c[cmd.name] = cmd.callback
    return c
}
