package main

import (
	"fmt"
	"os"
	"syscall"

	"golang.org/x/term"
)

func clear() {
	fmt.Print("\033[H\033[J")
}

func swapColors() {
	fmt.Print("\033[7m")
}

func resetColors() {
	fmt.Print("\033[0m")
}

func enterAlternativeScreen() {
	fmt.Print("\033[?1049h")
	fmt.Print("\033[?25l")
}

func exitAlternativeScreen() {
	fmt.Print("\033[?1049l")
	fmt.Print("\033[?25h")
}

func size() (int, int, error) {
	return term.GetSize(int(os.Stdout.Fd()))
}

func readKey() (rune, error) {
	oldState, err := term.MakeRaw(int(syscall.Stdin))
	if err != nil {
		return 0, err
	}
	defer term.Restore(int(syscall.Stdin), oldState)

	buf := make([]byte, 3)
	n, err := os.Stdin.Read(buf)
	if err != nil {
		return 0, err
	}

	// Detect arrow keys
	if n == 3 && buf[0] == 27 && buf[1] == 91 { // Escape sequence for arrow keys
		switch buf[2] {
		case 65:
			return '↑', nil
		case 66:
			return '↓', nil
		case 67:
			return '→', nil
		case 68:
			return '←', nil
		}

	}
	if n == 1 && buf[0] == 13 {
		return '\n', nil // Enter key
	}

	return rune(buf[0]), nil
}
