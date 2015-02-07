package sysexits

import "os"

// Exit exits the program (via os.Exit) with the provided status code.
func Exit(c ExitCode) {
	os.Exit(int(c))
}
