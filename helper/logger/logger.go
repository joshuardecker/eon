package logger

import "fmt"

// A logger is simply used for clarification on what systems are printing to console.
type Logger string

// Creates and returns a Logger with the given id.
func NewLogger(id string) *Logger {

	l := new(Logger)
	*l = Logger("[" + id + "]: ")

	return l
}

// Returns the string stored in Logger (the id of the logger).
func (l *Logger) String() string {

	return string(*l)
}

// Logs a message with a red id.
// Used for logging errors.
func (l *Logger) LogRed(msg string) {

	fmt.Println(Red(l.String()), msg)
}

// Logs a message with a yellow id.
// Used for logging update messages.
func (l *Logger) LogYellow(msg string) {

	fmt.Println(Yellow(l.String()), msg)
}

// Logs a message with a green id.
// Used for logging a valid event, like finding a block.
func (l *Logger) LogGreen(msg string) {

	fmt.Println(Green(l.String()), msg)
}

// Logs a message with a blue id.
// Used for logging the loading / creation of something, like a worker or the node loading the network.
func (l *Logger) LogBlue(msg string) {

	fmt.Println(Blue(l.String()), msg)
}
