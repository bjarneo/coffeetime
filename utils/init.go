package utils

import (
	"flag"
)

func Args() (string, int, int) {
	color := flag.String("color", "lightYellow", "the clock color")
	breakInterval := flag.Int("break-interval", 120, "The interval in minutes between coffee breaks")
	breakLength := flag.Int("break-length", 5, "The length in minutes of the coffee break")

	flag.Parse()

	return *color, *breakInterval, *breakLength
}
