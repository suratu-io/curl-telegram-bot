package commands

type (
    // Command simplifies work with commands by storing
    // command name and callback pair in a struct.
    Command struct {
        name     Name
        callback Callback
    }
)

// NewCommand returns a new Command.
func NewCommand(name Name, callback Callback) *Command {
    return &Command{name, callback}
}
