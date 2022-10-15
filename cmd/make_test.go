package main

import (
	"testing"
)

func TestCmdName(t *testing.T) {
	want := "gurl"
	got := cmdName()
	if got != want {
		t.Errorf("cmdName() = %q, want %q", got, want)
	}
}

func TestMain(t *testing.T) {
	main()
}
