package config

import (
    "github.com/suratu-io/curl-telegram-bot/app/logger"
    "os"
    "strconv"
)

const (
    BotApiTokenEnv    = "BOT_API_TOKEN"
    UpdateTimeoutEnv  = "UPDATE_TIMEOUT"
)

var (
    log = logger.Factory.NewLogger("config")

    BotApiToken    = os.Getenv(BotApiTokenEnv)
    UpdateTimeout  = 60
)

func init() {
    ut, err := strconv.ParseInt(os.Getenv(UpdateTimeoutEnv), 10, 32)
    if err != nil {
        log.Warn.Println("update timeout has not been set properly, using the default")
    } else {
        UpdateTimeout = int(ut)
    }
}
