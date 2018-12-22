package tty

/*
#include <stdio.h>
#if defined __APPLE__ || defined __linux__ || defined __unix__
#include <unistd.h>
	int IsTTY() { return isatty(fileno(stdout)) ? 1 : 0; }
#elif __WIN32 || _WIN64
#include <Windows.h>
	CONSOLE_SCREEN_BUFFER_INFO sbi;
	int IsTTY() { return GetConsoleScreenBufferInfo(GetStdHandle(STD_OUTPUT_HANDLE), &sbi) ? 1 : 0; }
#else
#endif
*/
import "C"

func itob(v C.int) bool { return v != 0 }
func IsTTY() bool       { return itob(C.IsTTY()) }
