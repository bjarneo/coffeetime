package main

import (
	"github.com/bjarneo/coffeetime/print"
	"github.com/bjarneo/coffeetime/utils"
)

func main() {
	color, breakInterval, breakLength := utils.Args()

	isBreak := utils.IsBreak(breakInterval, breakLength)

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
