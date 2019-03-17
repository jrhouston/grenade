package grenade

import "os"

// Crash causes a panic
func Crash() {
	panic("Something went horribly wrong.")
}

// Exit causes the process to exit with a non-zero exit code
func Exit() {
	os.Exit(137)
}
