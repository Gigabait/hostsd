package utils

import (
	// stdlib
	"flag"
	"os"
)

// ExitWhileNotInTests is a special helper that allows us to simulate FATAL
// logging level while using ERROR level. This is a workaroung for situations
// when we should exit immediately while running not in tests. If we're in
// tests - we will return false.
func ExitWhileNotInTests(exitcode int) bool {
	if flag.Lookup("test.v") == nil {
		os.Exit(exitcode)
	}

	return false
}
