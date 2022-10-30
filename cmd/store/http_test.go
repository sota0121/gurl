package store

import (
	"testing"
)

func TestGurlResponse_Status(t *testing.T) {
	want := "200 OK"
	got := GurlResponse{Status: want}.Status
	if got != want {
		t.Errorf("GurlResponse.Status = %q, want %q", got, want)
	}
}

func TestGurlResponse_Headers(t *testing.T) {
	want := map[string]string{"Content-Type": "text/plain"}
	got := GurlResponse{Headers: want}.Headers
	if got["Content-Type"] != want["Content-Type"] {
		t.Errorf("GurlResponse.Headers = %v, want %v", got, want)
	}
}
