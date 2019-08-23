package commander

// Command represents single command handler - second part of command.
// All other arguments besides first two will be sent to handler in
// parameters slice.
type Command struct {
	// Name is a command name.
	Name string
	// Description is a command description. Shown as help.
	Description string
	// Handler is a command handler.
	Handler func(parameters []string)
}
