package utils

var shouldRun bool = false

func ShouldRunOnce(isBreak bool) bool {
	if shouldRun && isBreak {
		return false
	}

	if !shouldRun && isBreak {
		shouldRun = true
	}

	if shouldRun && !isBreak {
		shouldRun = false
	}

	return shouldRun
}
