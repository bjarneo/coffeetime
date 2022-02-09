package main

import (
	"github.com/bjarneo/coffeetime/print"
	"github.com/bjarneo/coffeetime/utils"
)

func main() {
	flags := utils.Args()

	isBreak := utils.IsBreak(*flags.Interval, *flags.Duration)

	for {
		utils.Clear()

		print.Title()

		print.Clock(*flags.Color)

		if isBreak() {
			print.Coffee()
		}

		utils.Sleep()
	}
}
