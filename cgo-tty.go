package tty

/*
#include <stdio.h>
#include <sys/termios.h>
#if defined __APPLE__ || defined __linux__ || defined __unix__
#include <unistd.h>
	int IsTTY() {printf("Linux\n"); return isatty(fileno(stdout)) ? 1 : 0; }
#elif __WIN32 || _WIN64
#include <Windows.h>
	int isatty(int fd) {
		struct termios ts;
		return(tcgetattr(fd, &ts) != -1)
	}
	CONSOLE_SCREEN_BUFFER_INFO sbi;
	int IsTTY() {
		printf("windows\n");
		if (GetConsoleScreenBufferInfo(GetStdHandle(STD_OUTPUT_HANDLE), &sbi)) {
			printf("console\n");
			return 1;
		} else if (isatty(1)) {
			printf("tty\n");
			return 1;
		} else {
			printf("non\n");
			return 0;
		}
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
