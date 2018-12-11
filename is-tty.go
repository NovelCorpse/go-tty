package tty

import (
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

func IsTTY() bool { return terminal.IsTerminal(int(os.Stdout.Fd())) }
