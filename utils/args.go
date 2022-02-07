package utils

import (
	"flag"
)

func Args() (string, int, int) {
	color := flag.String("color", "lightYellow", "the clock color")
	interval := flag.Int("interval", 120, "The interval in minutes between coffee breaks")
	duration := flag.Int("duration", 5, "The duration in minutes of the coffee break")

	flag.Parse()

	return *color, *interval, *duration
}
