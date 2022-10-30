package feature

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGurlCommand_Execute(t *testing.T) {
	c := NewCommand()
	require.Equal(t, ReqContext{method: http.MethodGet}, *c.ctx)
	require.Equal(t, CmdConfig{}, *c.cfg)
	require.Equal(t, "dummy", c.client)
	require.Equal(t, "Usage: gurl [options...] <url>", c.usage, "usage should be equal")

	err := c.Execute()
	require.NoError(t, err)
}
