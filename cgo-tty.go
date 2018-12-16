package tty

/*
#include <stdio.h>
#if defined __APPLE__ || defined __linux__
#include <unistd.h>
	int IsTTY() { return isatty(fileno(stdout)) ? 1 : 0; }
#elif __WIN32 || _WIN64
#include <stdio.h>
#include <Windows.h>
	CONSOLE_SCREEN_BUFFER_INFO sbi;
	int IsTTY() { return GetConsoleScreenBufferInfo(GetStdHandle(STD_OUTPUT_HANDLE), &sbi) ? 1 : 0; }
#else
#endif
*/
import "C"

func IsTTY() bool {
	if C.IsTTY() == 1 {
		return true
	} else {
		return false
	}
}
