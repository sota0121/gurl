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
	for i := 0; i < GetOptSize(); i++ {
		opt, err := GetOptWithIndex(i)
		if err != nil {
			return err
		}
		fmt.Println(opt.GetLineToPrint())
		prefix := "value: "
		switch opt.GetType() {
		case String:
			str, err := opt.String()
			if err != nil {
				return err
			}
			fmt.Println(prefix, str)
		case Bool:
			b, err := opt.Bool()
			if err != nil {
				return err
			}
			fmt.Println(prefix, b)
		case Int:
			i, err := opt.Int()
			if err != nil {
				return err
			}
			fmt.Println(prefix, i)
		}
		fmt.Println()
	}
	return nil
}

// NewCommand creates a new command.
func NewCommand() *GurlCommand {
	ParserInit()
	return &GurlCommand{usage: "Usage: gurl [options...] <url>"}
}
