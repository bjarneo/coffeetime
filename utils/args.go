package utils

import (
	"flag"
)

type arguments struct {
	Color    *string
	Webhook  *string
	Interval *int
	Duration *int
	Notify   *bool
}

func Args() arguments {
	flags := arguments{
		Color:    flag.String("color", "lightYellow", "the clock color"),
		Webhook:  flag.String("webhook", "", "A webhook to call when your break begins. GET request."),
		Interval: flag.Int("interval", 120, "The interval in minutes between coffee breaks"),
		Duration: flag.Int("duration", 5, "The duration in minutes of the coffee break"),
		Notify:   flag.Bool("notify", false, "Notify the user with a system sound when it is time for a break"),
	}

	flag.Parse()

	return flags
}
