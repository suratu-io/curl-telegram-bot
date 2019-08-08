package commands

import (
    tbot "github.com/go-telegram-bot-api/telegram-bot-api"
)

type (
    // Name represents is the string value that represents
    // telegram bot command name.
    Name        string
    // Callback is the function that will be executed when
    // the appropriate command is fired.
    Callback    func(bot *tbot.BotAPI, msg tbot.Update) error
    // CallbackMap makes command selection easier.
    CallbackMap map[Name]Callback
)

// NewCallbackMap maps Commands into a CallbackMap.
func NewCallbackMap(commands ...*Command) CallbackMap {
    commandMap := CallbackMap{}
    for _, command := range commands {
        commandMap.SetFromCommand(command)
    }
    return commandMap
}

// Get returns a value from CallbackMap by its command name.
func (c CallbackMap) Get(name Name) (callback Callback, ok bool) {
    callback, ok = c[name]
    return
}

// Set sets a new callback by command name into the CallbackMap.
func (c CallbackMap) Set(name Name, callback Callback) CallbackMap {
    c[name] = callback
    return c
}

// SetFromCommand sets a new callback by command name into the CallbackMap
// taking values from the given Command.
func (c CallbackMap) SetFromCommand(cmd *Command) CallbackMap {
    c[cmd.name] = cmd.callback
    return c
}
