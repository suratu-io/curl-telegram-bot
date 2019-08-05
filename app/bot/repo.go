package bot

type (
    RunBotRepo interface {
        Run(timeout int) error
    }
)
