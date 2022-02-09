package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bjarneo/coffeetime/print"
	"github.com/bjarneo/coffeetime/request"
	"github.com/bjarneo/coffeetime/utils"
)

func main() {
	cancelChan := make(chan os.Signal, 1)
	signal.Notify(cancelChan, syscall.SIGTERM, syscall.SIGINT)

	go func() {
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
	}()

	<-cancelChan
	fmt.Printf("Aww, see you around")
}
