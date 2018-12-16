package tty

/*
#ifdef __linux__
	#include <stdio.h>
	#include <unistd.h>

	int IsTTY() {
		printf("linux tty \n");
    	return isatty(fileno(stdout)) ? 1 : 0;
    }
#elif _WIN32
    #include <stdio.h>
	#include <Windows.h>

	CONSOLE_SCREEN_BUFFER_INFO sbi;
	int IsTTY() {
    	return GetConsoleScreenBufferInfo(GetStdHandle(STD_OUTPUT_HANDLE), &sbi) ? 1 : 0;
	}
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
