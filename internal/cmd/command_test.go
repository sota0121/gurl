package cmd

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGurlCommand_Execute(t *testing.T) {
	c := NewCommand()
	require.Equal(t, "dummy", c.ctx)
	require.Equal(t, "dummy", c.cfg)
	require.Equal(t, "dummy", c.client)
	require.Equal(t, "Usage: gurl [options...] <url>", c.usage, "usage should be equal")

	err := c.Execute()
	require.NoError(t, err)
}
