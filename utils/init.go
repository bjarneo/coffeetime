package utils

import (
	"flag"
)

func Args() (string, int) {
	color := flag.String("color", "lightYellow", "the clock color")
	interval := flag.Int("interval", 120, "The interval in minutes between coffee breaks")

	flag.Parse()

	return *color, *interval
}
