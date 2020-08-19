package client

import (
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	timeout := time.Duration(5 * time.Second)
	client := NewClient(timeout)
	if client.Timeout != timeout {
		t.Error()
	}
}
