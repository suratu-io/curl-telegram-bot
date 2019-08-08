package bot

type (
    // RunBotRepo handles bots update chan.
    RunBotRepo interface {
        // Run starts listening to bot api and waits for new messages.
        Run(timeout int) error
    }
)
