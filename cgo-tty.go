package tty

/*
#include <stdio.h>
#include <unistd.h>

int IsTTY() {
    return isatty(fileno(stdout)) ? 1 : 0;
}
*/
import "C"

func IsTTY() bool {
	if C.IsTTY() == 1 {
		return true
	} else {
		return false
	}
}
