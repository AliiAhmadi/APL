package color

import "runtime"

var Active bool = true

const (
	RESET  = "\033[0m"
	RED    = "\033[31m"
	GREEN  = "\033[32m"
	YELLOW = "\033[33m"
	BLUE   = "\033[34m"
	PURPLE = "\033[35m"
	CYAN   = "\033[36m"
	GRAY   = "\033[37m"
	WHITE  = "\033[97m"
)

func init() {
	if runtime.GOOS == "windows" {
		Active = false
	}
}

func Red(text string) string {
	if Active {
		return RED + text + RESET
	}
	return text
}

func Green(text string) string {
	if Active {
		return GREEN + text + RESET
	}
	return text
}

func Yellow(text string) string {
	if Active {
		return YELLOW + text + RESET
	}
	return text
}

func Blue(text string) string {
	if Active {
		return BLUE + text + RESET
	}
	return text
}

func Purple(text string) string {
	if Active {
		return PURPLE + text + RESET
	}
	return text
}

func Cyan(text string) string {
	if Active {
		return CYAN + text + RESET
	}
	return text
}

func Gray(text string) string {
	if Active {
		return GRAY + text + RESET
	}
	return text
}

func White(text string) string {
	if Active {
		return WHITE + text + RESET
	}
	return text
}
