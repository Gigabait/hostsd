package logger

import (
	// stdlib
	"fmt"
	"os"
	"strings"
	"time"

	// other
	"github.com/rs/zerolog"
)

var (
	Logger            zerolog.Logger
	loggerInitialized bool
)

// Initialize initializes zerolog with proper formatting and log level.
func Initialize() {
	// Check environment for logger level.
	// Defaulting to INFO.
	loggerLevel, loggerLevelFound := os.LookupEnv("LOGGER_LEVEL")
	if loggerLevelFound {
		switch strings.ToUpper(loggerLevel) {
		case "DEBUG":
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
		case "INFO":
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
		case "WARN":
			zerolog.SetGlobalLevel(zerolog.WarnLevel)
		case "ERROR":
			zerolog.SetGlobalLevel(zerolog.ErrorLevel)
		case "FATAL":
			zerolog.SetGlobalLevel(zerolog.FatalLevel)
		}

	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	output := zerolog.ConsoleWriter{Out: os.Stdout, NoColor: false, TimeFormat: time.RFC3339}
	output.FormatLevel = func(i interface{}) string {
		var v string
		if ii, ok := i.(string); ok {
			ii = strings.ToUpper(ii)
			switch ii {
			case "DEBUG":
				v = fmt.Sprintf("\x1b[30m%-5s\x1b[0m", ii)
			case "ERROR":
				v = fmt.Sprintf("\x1b[31m%-5s\x1b[0m", ii)
			case "FATAL":
				v = fmt.Sprintf("\x1b[35m%-5s\x1b[0m", ii)
			case "INFO":
				v = fmt.Sprintf("\x1b[32m%-5s\x1b[0m", ii)
			case "PANIC":
				v = fmt.Sprintf("\x1b[36m%-5s\x1b[0m", ii)
			case "WARN":
				v = fmt.Sprintf("\x1b[33m%-5s\x1b[0m", ii)
			default:
				v = ii
			}
		}
		return fmt.Sprintf("| %s |", v)
	}

	Logger = zerolog.New(output).With().Timestamp().Logger()
	loggerInitialized = true
}
