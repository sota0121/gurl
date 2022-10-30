package proxy

import (
	"testing"
)

func TestHttpClient_Connect(t *testing.T) {
	c := NewHttpClient()
	if c.IsConnected() {
		t.Errorf("HttpClient.IsConnected() = %v, want %v", c.IsConnected(), false)
	}
	if err := c.Connect(); err != nil {
		t.Errorf("HttpClient.Connect() = %v, want %v", err, nil)
	}
	if !c.IsConnected() {
		t.Errorf("HttpClient.IsConnected() = %v, want %v", c.IsConnected(), true)
	}
}

func TestHttpClient_IsConnected(t *testing.T) {
	c := NewHttpClient()
	if c.IsConnected() {
		t.Errorf("HttpClient.IsConnected() = %v, want %v", c.IsConnected(), false)
	}
}
