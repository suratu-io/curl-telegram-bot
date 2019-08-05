package commands

type (
    Command struct {
        name     Name
        callback Callback
    }
)

func NewCommand(name Name, callback Callback) *Command {
    return &Command{name, callback}
}
