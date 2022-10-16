package cmd

import (
	"flag"
	"fmt"
)

// ArgValueType is the type of the argument value.
type ArgValueType string

const (
	String ArgValueType = "string"
	Bool   ArgValueType = "bool"
	Int    ArgValueType = "int"
)

// OptValue is the value of the option with flexible type.
type OptValue struct {
	typeOfOpt ArgValueType
	strValue  string
	boolValue bool
	intValue  int
}

func NewOptValue(typeOfOpt ArgValueType) *OptValue {
	// create a new OptValue
	return &OptValue{
		typeOfOpt: typeOfOpt,
		strValue:  "",
		boolValue: false,
		intValue:  0,
	}
}

// func (v *OptValue) Set(value string) error {
// 	switch v.typeOfOpt {
// 	case String:
// 		v.strValue = value
// 	case Bool:
// 		v.boolValue = true
// 	case Int:
// 		intValue, err := strconv.Atoi(value)
// 		if err != nil {
// 			return err
// 		}
// 		v.intValue = intValue
// 	}
// 	return nil
// }

func (v OptValue) GetType() ArgValueType {
	return v.typeOfOpt
}

func (v OptValue) String() (string, error) {
	return v.strValue, nil
}

func (v OptValue) Bool() (bool, error) {
	return v.boolValue, nil
}

func (v OptValue) Int() (int, error) {
	return v.intValue, nil
}

// Opt is the container for the option.
type Opt struct {
	Name      OptNameType
	ShortHand string
	Help      string
	Usage     string
	Required  bool
	Supported bool
	Value     OptValue
}

func (o Opt) GetLineToPrint() string {
	supportStatus := "supported"
	if !o.Supported {
		supportStatus = "not supported"
	}
	return fmt.Sprintf(" %s,\t%s\t%s (%s)", o.ShortHand, o.Usage, o.Help, supportStatus)
}

func (o Opt) GetType() ArgValueType {
	return o.Value.GetType()
}

func (o Opt) String() (string, error) {
	return o.Value.String()
}

func (o Opt) Bool() (bool, error) {
	return o.Value.Bool()
}

func (o Opt) Int() (int, error) {
	return o.Value.Int()
}

func (o Opt) GetVal() OptValue {
	return o.Value
}

// OptNameType is the type of the option name.
type OptNameType string

func (o OptNameType) String() string {
	return string(o)
}

// Option Names definition
const (
	data       OptNameType = "data"
	fail       OptNameType = "fail"
	help       OptNameType = "help"
	include    OptNameType = "include"
	output     OptNameType = "output"
	remoteName OptNameType = "remote-name"
	silent     OptNameType = "silent"
	uploadFile OptNameType = "upload-file"
	user       OptNameType = "user"
	userAgent  OptNameType = "user-agent"
	verbose    OptNameType = "verbose"
	version    OptNameType = "version"
)

// options is the slice of the options.
var options = []Opt{
	{
		Name:      data,
		ShortHand: "-d",
		Help:      "HTTP POST data",
		Usage:     "--data <data>",
		Required:  false,
		Supported: true,
		Value:     *NewOptValue(String),
	},
	{
		Name:      fail,
		ShortHand: "-f",
		Help:      "Fail silently (no output at all) on HTTP errors",
		Usage:     "--fail",
		Required:  false,
		Supported: true,
		Value:     *NewOptValue(Bool),
	},
	{
		Name:      help,
		ShortHand: "-h",
		Help:      "Show this help message and exit",
		Usage:     "--help <category>",
		Required:  false,
		Supported: true,
		Value:     *NewOptValue(Bool),
	},
	{
		Name:      include,
		ShortHand: "-i",
		Help:      "Include protocol headers in the output (H/F)",
		Usage:     "--include",
		Required:  false,
		Supported: true,
		Value:     *NewOptValue(Bool),
	},
	{
		Name:      output,
		ShortHand: "-o",
		Help:      "Write to <file> instead of stdout",
		Usage:     "--output <file>",
		Required:  false,
		Supported: true,
		Value:     *NewOptValue(String),
	},
	{
		Name:      remoteName,
		ShortHand: "-O",
		Help:      "Write output to a file named as the remote file",
		Usage:     "--remote-name",
		Required:  false,
		Supported: true,
		Value:     *NewOptValue(Bool),
	},
	{
		Name:      silent,
		ShortHand: "-s",
		Help:      "Silent mode (don't output anything)",
		Usage:     "--silent",
		Required:  false,
		Supported: true,
		Value:     *NewOptValue(Bool),
	},
	{
		Name:      uploadFile,
		ShortHand: "-T",
		Help:      "Transfer local FILE to destination",
		Usage:     "--upload-file <file>",
		Required:  false,
		Supported: true,
		Value:     *NewOptValue(String),
	},
	{
		Name:      user,
		ShortHand: "-u",
		Help:      "Server user and password",
		Usage:     "--user <user:password>",
		Required:  false,
		Supported: true,
		Value:     *NewOptValue(String),
	},
	{
		Name:      userAgent,
		ShortHand: "-A",
		Help:      "User-Agent to send to server",
		Usage:     "--user-agent <name>",
		Required:  false,
		Supported: true,
		Value:     *NewOptValue(String),
	},
	{
		Name:      verbose,
		ShortHand: "-v",
		Help:      "Make the operation more talkative",
		Usage:     "--verbose",
		Required:  false,
		Supported: true,
		Value:     *NewOptValue(Bool),
	},
	{
		Name:      version,
		ShortHand: "-V",
		Help:      "Show version number and quit",
		Usage:     "--version",
		Required:  false,
		Supported: true,
		Value:     *NewOptValue(Bool),
	},
}

// Init initializes the parser.
func ParserInit() {
	for i, opt := range options {
		// NOTE: We use the index of the option in the slice as the flag name.
		// If we don't access the flag by index, the flag value will not be updated.
		// This is Go language's behavior.
		// In order to update array elements, we need to use the index.
		if opt.Value.GetType() == String {
			flag.StringVar(&options[i].Value.strValue, opt.Name.String(), "", opt.GetLineToPrint())
		} else if opt.Value.GetType() == Bool {
			flag.BoolVar(&options[i].Value.boolValue, opt.Name.String(), false, opt.GetLineToPrint())
		} else if opt.Value.GetType() == Int {
			flag.IntVar(&options[i].Value.intValue, opt.Name.String(), 0, opt.GetLineToPrint())
		}
	}
	flag.Parse()
}

// GetOptSize returns the number of options.
func GetOptSize() int {
	return len(options)
}

// GetOptWithIndex returns the option at the given index.
func GetOptWithIndex(index int) (Opt, error) {
	if index < 0 || index >= len(options) {
		return Opt{}, fmt.Errorf("index out of range: %d", index)
	}
	return options[index], nil
}

// GetOptWithName returns the option with the given name.
func GetOptWithName(name OptNameType) (Opt, error) {
	for _, opt := range options {
		if opt.Name == name {
			return opt, nil
		}
	}
	return Opt{}, fmt.Errorf("option not found: %s", name)
}
