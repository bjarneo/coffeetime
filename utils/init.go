package utils

import "time"

func Clear() {
	print("\033[H\033[2J")
}

func Sleep() {
	time.Sleep(time.Duration(1) * time.Second)
}
