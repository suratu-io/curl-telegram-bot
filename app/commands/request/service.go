package request

import (
    "fmt"
    tbot "github.com/go-telegram-bot-api/telegram-bot-api"
    "github.com/suratu-io/curl-telegram-bot/app/commands/request/arguments"
    "io/ioutil"
    "net/http"
    "strings"
)

// PerformRequest creates and sends a new http request configured by the set
// of arguments in the message.
func PerformRequest(method string, bot *tbot.BotAPI, msg tbot.Update) error {
    commandArgs := msg.Message.CommandArguments()
    splitArgs := strings.Split(commandArgs, " ")

    log.Debug.PrintTKV("request command - {{method}}, args: {{args}}",
        "method", method,
        "args", splitArgs,
    )

    args, err := arguments.ParseArguments(splitArgs...)
    if err != nil {
        return err
    }

    req, err := http.NewRequest(method, args.URL.String(), nil)
    if err != nil {
        return err
    }

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return err
    }

    log.Debug.Println(*resp)

    text := fmt.Sprintf(
        "Status: %v",
        resp.Status,
    )

    bodyBytes, err := ioutil.ReadAll(resp.Body)
    if err == nil {
        text += "\n" + string(bodyBytes)
    }

    message := tbot.NewMessage(msg.Message.Chat.ID, text)

    _, err = bot.Send(message)

    return err
}
