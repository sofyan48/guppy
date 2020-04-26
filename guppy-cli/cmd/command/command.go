package command

// Command ...
type Command struct {
}

// CommandHandler ...
func CommandHandler() *Command {
	return &Command{}
}

// CommandInterface ...
type CommandInterface interface{}
