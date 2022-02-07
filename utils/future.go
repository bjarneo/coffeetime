package utils

import (
	"fmt"
	"time"
)

func futureTime(interval int) time.Time {
	t := time.Now()

	future := t.Add(time.Minute * time.Duration(interval))

	return future
}

/*
func next(time time.Time) func() time.Time {
	var nextTime time.Time = time

	return func() time.Time {

	}
}
*/
func IsBreak(interval int, breakLength int) func() bool {
	var timeFormat string = "15:04:05 - 01-02-2006"

	future := futureTime(interval).Format(timeFormat)

	futureBreak := futureTime(breakLength)

	return func() bool {
		fmt.Println(futureBreak)
		fmt.Println(futureBreak.UnixNano() / 1000)
		now := time.Now().Format(timeFormat)

		if now == future {
			return true
		}

		return false
	}
}
