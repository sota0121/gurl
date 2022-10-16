package cmd

import (
	"testing"
)

func TestGurlCommand_Execute(t *testing.T) {
	c := &GurlCommand{}
	if err := c.Execute(); err != nil {
		t.Error(err)
	}
}
