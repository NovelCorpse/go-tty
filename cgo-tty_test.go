package tty

import (
	"fmt"
	"testing"
)

func TestIsTTY(t *testing.T) {
	fmt.Printf("Is this device a tty? - %t\n", IsTTY())
	fmt.Printf("TTY Size: \n")
	TermSize()
}
