package cmd

import (
	"fmt"
)

// GurlCommand is the main command.
type GurlCommand struct {
	ctx string // dummy field
	cfg string // dummy field

	client string // dummy field
	usage  string // dummy field
}

func (c *GurlCommand) Execute() error {
	fmt.Println(c.usage)
	fmt.Println("Execute called on GurlCommand")
	fmt.Println("-----------------------------")
	parser := NewArgParser()
	for i := 0; i < parser.GetOptSize(); i++ {
		opt, err := parser.GetOptWithIndex(i)
		if err != nil {
			return err
		}
		fmt.Println(opt.GetLineToPrint())
	}
	return nil
}

// NewCommand creates a new command.
func NewCommand() *GurlCommand {
	// TODO: parse args
	return &GurlCommand{usage: "Usage: gurl [options...] <url>"}
}
