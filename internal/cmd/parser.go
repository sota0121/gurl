package cmd

import (
	"flag"
	"fmt"
)

type Opt struct {
	Name      string
	ShortHand string
	Help      string
	Usage     string
	Required  bool
	Supported bool
	Value     string
}

func (o Opt) GetLineToPrint() string {
	return fmt.Sprintf(" %s,\t%s\t%s", o.ShortHand, o.Usage, o.Help)
}

var options = []Opt{
	{
		Name:      "data",
		ShortHand: "-d",
		Help:      "HTTP POST data",
		Usage:     "--data <data>",
		Required:  false,
		Supported: true,
	},
	{
		Name:      "fail",
		ShortHand: "-f",
		Help:      "Fail silently (no output at all) on HTTP errors",
		Usage:     "--fail",
		Required:  false,
		Supported: true,
	},
	{
		Name:      "help",
		ShortHand: "-h",
		Help:      "Show this help message and exit",
		Usage:     "--help <category>",
		Required:  false,
		Supported: true,
	},
	{
		Name:      "include",
		ShortHand: "-i",
		Help:      "Include protocol headers in the output (H/F)",
		Usage:     "--include",
		Required:  false,
		Supported: true,
	},
	{
		Name:      "output",
		ShortHand: "-o",
		Help:      "Write to <file> instead of stdout",
		Usage:     "--output <file>",
		Required:  false,
		Supported: true,
	},
	{
		Name:      "remote-name",
		ShortHand: "-O",
		Help:      "Write output to a file named as the remote file",
		Usage:     "--remote-name",
		Required:  false,
		Supported: true,
	},
	{
		Name:      "silent",
		ShortHand: "-s",
		Help:      "Silent mode (don't output anything)",
		Usage:     "--silent",
		Required:  false,
		Supported: true,
	},
	{
		Name:      "upload-file",
		ShortHand: "-T",
		Help:      "Transfer local FILE to destination",
		Usage:     "--upload-file <file>",
		Required:  false,
		Supported: true,
	},
	{
		Name:      "user",
		ShortHand: "-u",
		Help:      "Server user and password",
		Usage:     "--user <user:password>",
		Required:  false,
		Supported: true,
	},
	{
		Name:      "user-agent",
		ShortHand: "-A",
		Help:      "User-Agent to send to server",
		Usage:     "--user-agent <name>",
		Required:  false,
		Supported: true,
	},
	{
		Name:      "verbose",
		ShortHand: "-v",
		Help:      "Make the operation more talkative",
		Usage:     "--verbose",
		Required:  false,
		Supported: true,
	},
	{
		Name:      "version",
		ShortHand: "-V",
		Help:      "Show version number and quit",
		Usage:     "--version",
		Required:  false,
		Supported: true,
	},
}

// ArgParser is the parser for command line arguments.
type ArgParser struct {
	options []Opt
}

// NewArgParser creates a new parser.
func NewArgParser() *ArgParser {
	parser := &ArgParser{
		options: options,
	}
	parser.Init()
	return parser
}

// Init initializes the parser.
func (p *ArgParser) Init() {
	for _, opt := range p.options {
		flag.StringVar(&opt.Value, opt.Name, "", opt.GetLineToPrint())
	}
	flag.Parse()
}

// GetOptSize returns the number of options.
func (p *ArgParser) GetOptSize() int {
	return len(p.options)
}

// GetOptWithIndex returns the option at the given index.
func (p *ArgParser) GetOptWithIndex(index int) (Opt, error) {
	if index < 0 || index >= len(p.options) {
		return Opt{}, fmt.Errorf("index out of range: %d", index)
	}
	return p.options[index], nil
}

// GetOptWithName returns the option with the given name.
func (p *ArgParser) GetOptWithName(name string) (Opt, error) {
	for _, opt := range p.options {
		if opt.Name == name {
			return opt, nil
		}
	}
	return Opt{}, fmt.Errorf("option not found: %s", name)
}
