package helper

import (
	"time"
)

// Return the unix time as a int64.
func UnixTime() int64 {

	return time.Now().Unix()
}

// Return the local time in the time.Time format.
func LocalTime() time.Time {

	return time.Now()
}

// ****
// Timer:

// A Timer starts as a int64 that holds the start unix time.
type Timer int64

// Return a new timer, bound to the unix time of its creation.
func CreateTimer() Timer {

	return Timer(UnixTime())
}

// Get the unix time stored in the timer.
func (t *Timer) Time() int64 {

	return int64(*t)
}

// Get the time (seconds) that have passed since the timer was created.
func (t *Timer) GetTime() int64 {

	return (UnixTime() - t.Time())
}

// Reset the time stored in the timer.
func (t *Timer) Reset() {

	*t = Timer(UnixTime())
}

// Timer:
// ****
