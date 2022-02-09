package main

import (
	"github.com/bjarneo/coffeetime/print"
	"github.com/bjarneo/coffeetime/request"
	"github.com/bjarneo/coffeetime/utils"
)

func main() {
	flags := utils.Args()

	isBreak := utils.IsBreak(*flags.Interval, *flags.Duration)

	for {
		utils.Clear()

		print.Title()

		print.Clock(*flags.Color)

		if utils.ShouldRunOnce(isBreak()) {
			request.GetHook(*flags.Webhook)
		}

		if isBreak() {
			print.Coffee()
		}

		utils.Sleep()
	}
}
