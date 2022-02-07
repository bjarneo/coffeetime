package utils

import (
	"github.com/pterm/pterm"
)

func Color(color string) pterm.Color {
	colors := map[string]pterm.Color{
		"red":     pterm.FgRed,
		"cyan":    pterm.FgCyan,
		"gray":    pterm.FgGray,
		"blue":    pterm.FgBlue,
		"black":   pterm.FgBlack,
		"green":   pterm.FgGreen,
		"white":   pterm.FgWhite,
		"yellow":  pterm.FgYellow,
		"magenta": pterm.FgMagenta,

		"lightRed":     pterm.FgLightRed,
		"lightCyan":    pterm.FgLightCyan,
		"lightBlue":    pterm.FgLightBlue,
		"lightGreen":   pterm.FgLightGreen,
		"lightWhite":   pterm.FgLightWhite,
		"lightYellow":  pterm.FgLightYellow,
		"lightMagenta": pterm.FgLightMagenta,
	}

	value, ok := colors[color]

	// If the color from the cli is used and is not OK, return light yellow.
	if !ok {
		return colors["lightYellow"]
	}

	return value
}
