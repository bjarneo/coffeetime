package main

import (
	"github.com/bjarneo/coffeetime/print"
	"github.com/bjarneo/coffeetime/utils"
)

func main() {
	color, interval, duration := utils.Args()

	isBreak := utils.IsBreak(interval, duration)

	for {
		utils.Clear()

		print.Title()

		print.Clock(color)

		if isBreak() {
			print.Coffee()
		}

		utils.Sleep()
	}
}
