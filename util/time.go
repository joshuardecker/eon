package utilities

import (
	"time"
)

type Time struct {
	timerLog   uint64
	timerValue uint64
}

// Function returns the current unix time in seconds.
func (t *Time) CurrentUnix() uint64 {

	return uint64(time.Now().Unix())
}

// Function returns the current unix time in milli-seconds.
func (t *Time) CurrentUnixMilli() uint64 {

	return uint64(time.Now().UnixMicro())
}

// Returns a nice version of the local time.
func (t *Time) CurrentTime() time.Time {

	return time.Now().Round(time.Second)
}

func (t *Time) Timer() uint64 {

	t.timerLog = t.timerValue

	t.timerValue = t.CurrentUnix()

	return t.timerValue - t.timerLog
}
