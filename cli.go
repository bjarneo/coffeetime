package main

import (
	"fmt"
	"time"

	"github.com/bjarneo/coffeetime/utils"
	"github.com/pterm/pterm"
)

func clear() {
	print("\033[H\033[2J")
}

func printTitle() {
	now := time.Now()

	fmt.Println("\033]0; " + now.Format("15:04:05 - 01-02-2006") + " \007")
}

func printClock(color string) {
	now := time.Now()

	timeString, _ := pterm.DefaultBigText.WithLetters(
		pterm.NewLettersFromStringWithStyle(
			now.Format("15:04:05"),
			pterm.NewStyle(utils.Color(color)),
		),
	).Srender()

	pterm.DefaultCenter.Println("\n\n")

	pterm.DefaultCenter.Println(timeString)

	pterm.DefaultCenter.Println(now.Format("2006-01-02"))
}

var coffeeIncrement int = 0

func printCoffee() {
	var coffee string = `
    ( (
   ) )
 ........
 |      |]
 \      /
  \____/`

	var coffee2 string = `
    ) )
   ( (
 ........
 |      |]
 \      /
  \____/`

	if coffeeIncrement%2 == 0 {
		pterm.DefaultCenter.Println(coffee)
	} else {
		pterm.DefaultCenter.Println(coffee2)
	}

	pterm.DefaultCenter.Println("It is time for a coffee break")

	coffeeIncrement++
}

func run() {
	color, breakInterval, breakLength := utils.Args()

	isBreak := utils.IsBreak(breakInterval, breakLength)

	for {
		clear()

		printTitle()

		printClock(color)

		if isBreak() {
			printCoffee()
		}

		time.Sleep(time.Duration(1) * time.Second)
	}
}

func main() {
	run()
}
