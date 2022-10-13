package logger

// Console Colors:
const (
	DEFAULT = "\x1b[39m"
	RED     = "\x1b[31m"
	YELLOW  = "\x1b[33m"
	GREEN   = "\x1b[32m"
	BLUE    = "\x1b[36m"
)

// Changes the str color to red.
func Red(str string) (colorStr string) {

	return RED + str + DEFAULT
}

// Changes the str color to yellow.
func Yellow(str string) (colorStr string) {

	return YELLOW + str + DEFAULT
}

// Changes the str color to green.
func Green(str string) (colorStr string) {

	return GREEN + str + DEFAULT
}

// Changes the str color to blue.
func Blue(str string) (colorStr string) {

	return BLUE + str + DEFAULT
}
