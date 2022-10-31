package feature

import (
	"flag"
	"net/http"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	progname = "gurl"
)

type IntOpt struct { // implements for flag.Value interface
	value int
	name  string
}

func (o *IntOpt) Set(value string) error {
	converted, err := strconv.Atoi(value)
	if err != nil {
		return err
	}
	o.value = converted
	return nil
}

func (o IntOpt) String() string {
	return strconv.Itoa(o.value)
}

type StringOpt struct { // implements for flag.Value interface
	value string
	name  string
}

func (o *StringOpt) Set(value string) error {
	o.value = value
	return nil
}

func (o StringOpt) String() string {
	return o.value
}

type BoolOpt struct { // implements for flag.Value interface
	value bool
	name  string
}

func (o *BoolOpt) Set(value string) error {
	converted, err := strconv.ParseBool(value)
	if err != nil {
		return err
	}
	o.value = converted
	return nil
}

func (o BoolOpt) String() string {
	return strconv.FormatBool(o.value)
}

// setCommandlineArgs sets the commandline arguments for testing
func setCommandlineArgs(t *testing.T, progname string, args []string) {
	t.Helper() // mark this function as a helper function

	os.Args = append([]string{progname}, args...)
}

// setCommandlineOptions sets the commandline options for testing
func setCommandlineOptions(t *testing.T, opts []flag.Value) {
	t.Helper() // mark this function as a helper function

	for _, opt := range opts {
		switch valuetype := opt.(type) {
		case *IntOpt:
			flag.CommandLine.Set(opt.(*IntOpt).name, opt.(*IntOpt).String())
		case *StringOpt:
			flag.CommandLine.Set(opt.(*StringOpt).name, opt.(*StringOpt).String())
		case *BoolOpt:
			flag.CommandLine.Set(opt.(*BoolOpt).name, opt.(*BoolOpt).String())
		default:
			t.Fatalf("unknown option type: %T", valuetype)
		}
	}
}

func TestNewCommand(t *testing.T) {
	// [Essence] At unittest, we should use the virtual commandline arguments
	// instead of the real commandline arguments.
	// [Note] The commandline arguments are stored in the os.Args.
	// [Note] The commandline options are stored with the flag.CommandLine.Set()
	// flag.CommandLine.Set() has to be called before flag.Parse()
	setCommandlineArgs(t, progname, []string{"GET", "http://example.com"})
	testoptions := make([]flag.Value, 0, 3)
	testoptions = append(testoptions, &StringOpt{value: "hello", name: "data"})
	testoptions = append(testoptions, &BoolOpt{value: true, name: "verbose"})
	setCommandlineOptions(t, testoptions)

	c := NewCommand()
	require.Equal(t, ReqContext{method: http.MethodGet}, *c.ctx)
	require.Equal(t, CmdConfig{}, *c.cfg)
	require.Equal(t, "dummy", c.client)
	require.Equal(t, "Usage: gurl [options...] <url>", c.usage, "usage should be equal")
}
