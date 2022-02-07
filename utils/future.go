package utils

import (
	"time"
)

func futureTime(breakInterval int) time.Time {
	t := time.Now()

	future := t.Add(time.Minute * time.Duration(breakInterval))

	return future
}

var timestamp = map[string]int64{
	"futureBreakStart": 0,
	"futureBreakStop":  0,
}

func updateTimestamp(breakInterval int, breakLength int) {
	timestamp["futureBreakStart"] = futureTime(breakInterval).UnixMilli()
	timestamp["futureBreakStop"] = futureTime(breakInterval + breakLength).UnixMilli()
}

func IsBreak(breakInterval int, breakLength int) func() bool {
	updateTimestamp(breakInterval, breakLength)

	return func() bool {
		now := time.Now().UnixMilli()

		if now >= timestamp["futureBreakStart"] && now <= timestamp["futureBreakStop"] {
			return true
		}

		if now >= timestamp["futureBreakStop"] {
			updateTimestamp(breakInterval, breakLength)
		}

		return false
	}
}
