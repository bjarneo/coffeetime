package utils

import (
	"github.com/gen2brain/beeep"
)

func Notify(notify bool) bool {
	if !notify {
		return false
	}

	err := beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)

	if err != nil {
		panic(err)
	}

	return true
}
