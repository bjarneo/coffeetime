package request

import (
	"net/http"
	"time"
)

var shouldRun bool = false

func httpClient() *http.Client {
	return &http.Client{Timeout: 10 * time.Second}
}

func GetHook(webhook string, isBreak bool) bool {
	if !shouldRun && isBreak {
		shouldRun = true
	}

	if webhook == "" || !shouldRun {
		return false
	}

	resp, err := http.Get(webhook)

	shouldRun = false

	if err != nil {
		return false
	}

	defer resp.Body.Close()

	return true
}
