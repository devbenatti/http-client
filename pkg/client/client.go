package client

import (
	"net/http"
	"time"
)

// newClient build http.Client
func NewClient(timeout time.Duration) http.Client {
	return http.Client{
		Timeout: timeout,
	}
}
