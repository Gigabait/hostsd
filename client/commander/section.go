package commander

// This structure describes command section - first part of command.
type section struct {
	// Name is a section name, like "config", "status", etc.
	Name string
	// Commands is a map with commands where key is a command name
	// and value is a Command struct which describes command.
	Commands map[string]*Command
}
