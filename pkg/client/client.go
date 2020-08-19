package client

import (
	"net/http"
	"time"
)

// NewClient build http.Client
func NewClient(timeout time.Duration) http.Client {
	return http.Client{
		Timeout: timeout,
	}
}
