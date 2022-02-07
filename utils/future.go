package utils

import (
	"time"
)

func futureTime(interval int) time.Time {
	t := time.Now()

	future := t.Add(time.Minute * time.Duration(interval))

	return future
}

var timestamp = map[string]int64{
	"futureBreakStart": 0,
	"futureBreakStop":  0,
}

func updateTimestamp(interval int, duration int) {
	timestamp["futureBreakStart"] = futureTime(interval).UnixMilli()
	timestamp["futureBreakStop"] = futureTime(interval + duration).UnixMilli()
}

func IsBreak(interval int, duration int) func() bool {
	updateTimestamp(interval, duration)

	return func() bool {
		now := time.Now().UnixMilli()

		if now >= timestamp["futureBreakStart"] && now <= timestamp["futureBreakStop"] {
			return true
		}

		if now >= timestamp["futureBreakStop"] {
			updateTimestamp(interval, duration)
		}

		return false
	}
}
