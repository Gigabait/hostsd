package commander

import (
	// stdlib
	"os"

	// local
	"github.com/medium-isp/hostsd/internal/logger"

	// other
	"github.com/rs/zerolog"
)

var (
	log               zerolog.Logger
	loggerInitialized bool

	commands map[string]*section
)

func AddCommand(sectionName string, command *Command) {
	_, sectionExists := commands[sectionName]
	if !sectionExists {
		commands[sectionName] = &section{
			Name:     sectionName,
			Commands: make(map[string]*Command),
		}
	}

	commands[sectionName].Commands[command.Name] = command
	log.Debug().Str("section", sectionName).Str("command", command.Name).Msg("Command registered")
}

func Initialize() {
	log = logger.Logger.With().Str("package", "commander").Logger()
	loggerInitialized = true

	log.Debug().Msg("Initializing...")
	commands = make(map[string]*section)
}

func Process() {
	// We should receive at least two parameters - section and command.
	if len(os.Args) < 3 {
		log.Fatal().Msg("hostsdctl need at least two parameters (section and command). See 'hostsdctl commander help' for list of available sections and commands.")
	}
	sectionName := os.Args[1]
	commandName := os.Args[2]
	params := os.Args[3:]
	log.Debug().Str("section", sectionName).Str("command", commandName).Interface("parameters", params).Msg("CLI parameters parsed")

	// Checking for section availability.
	section, sectionExists := commands[sectionName]
	if !sectionExists {
		log.Fatal().Str("section", sectionName).Msg("Section does not exist. See 'hostsdctl commander help' for list of available sections and commands.")
	}

	// Checking for command availability.
	command, commandExists := section.Commands[commandName]
	if !commandExists {
		log.Fatal().Str("command", commandName).Msg("Command does not exist. See 'hostsdctl commander help' for list of available sections and commands.")
	}

	// Execute command.
	command.Handler(params)
}
