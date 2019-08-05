package commands

import (
    "github.com/olehan/kek"
    "github.com/olehan/kek/formatters/sugared"
    "os"
)

var (
    Factory = kek.NewFactory(os.Stdout, sugared.Formatter, "commands")
)

func init() {
    Factory.SetRandomNameColor().SetWithNS(false)
}
