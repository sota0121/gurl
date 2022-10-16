package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGurlCommand_Execute(t *testing.T) {
	c := NewCommand()
	assert.Equal(t, "dummy", c.ctx)
	assert.Equal(t, "dummy", c.cfg)
	assert.Equal(t, "dummy", c.client)
	assert.Equal(t, "Usage: gurl [options...] <url>", c.usage, "usage should be equal")

	if err := c.Execute(); err != nil {
		assert.Fail(t, err.Error(), "GurlCommand.Execute() should not return error")
	}
}
