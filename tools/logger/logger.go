package logger

import "fmt"

// Logs a message with a red id.
// Used for logging errors.
func LogRed(sender, msg string) {

	fmt.Println(Red(sender)+": ", msg)
}

// Logs a message with a yellow id.
// Used for logging update messages.
func LogYellow(sender, msg string) {

	fmt.Println(Yellow(sender)+": ", msg)
}

// Logs a message with a green id.
// Used for logging a valid event, like finding a block.
func LogGreen(sender, msg string) {

	fmt.Println(Green(sender)+": ", msg)
}

// Logs a message with a blue id.
// Used for logging the loading / creation of something, like a worker or the node loading the network.
func LogBlue(sender, msg string) {

	fmt.Println(Blue(sender)+": ", msg)
}
