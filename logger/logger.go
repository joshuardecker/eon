package logger

import "fmt"

type Logger struct {
	id string
}

// Creates a new logger
func NewLogger(id string) *Logger {

	l := new(Logger)
	l.id = "[" + id + "]: "

	return l
}

// Logs a message with a red id.
// Used for logging errors.
func (l *Logger) LogRed(msg string) {

	fmt.Println(Red(l.id), msg)
}

// Logs a message with a yellow id.
// Used for logging update messages.
func (l *Logger) LogYellow(msg string) {

	fmt.Println(Yellow(l.id), msg)
}

// Logs a message with a green id.
// Used for logging a valid event, like finding a block.
func (l *Logger) LogGreen(msg string) {

	fmt.Println(Green(l.id), msg)
}

// Logs a message with a blue id.
// Used for logging the loading / creation of something, like a worker or the node loading the network.
func (l *Logger) LogBlue(msg string) {

	fmt.Println(Blue(l.id), msg)
}
