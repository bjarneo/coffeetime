package request

import (
	"net/http"
	"time"
)

var shouldRun bool = false

func httpClient() *http.Client {
	return &http.Client{Timeout: 15 * time.Second}
}

func GetHook(webhook string) bool {
	if webhook == "" {
		return false
	}

	res, err := httpClient().Get(webhook)

	if err != nil {
		return false
	}

	defer res.Body.Close()

	return true
}
