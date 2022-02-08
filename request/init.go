package request

import (
	"net/http"
	"time"
)

func httpClient() *http.Client {
	return &http.Client{Timeout: 10 * time.Second}
}

func Hook(webhook string) bool {
	if webhook == "" {
		return false
	}

	resp, err := http.Get(webhook)

	if err != nil {
		return false
	}

	defer resp.Body.Close()

	return true
}
