package feature

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

func (v OptValue) String() string {
	return v.strValue
}

func (v OptValue) Bool() bool {
	return v.boolValue
}

func (v OptValue) Int() int {
	return v.intValue
}

// Opt is the container for the option.
type Opt struct {
	Name      OptNameType
	ShortHand OptNameType
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

func (o Opt) String() string {
	return o.Value.String()
}

func (o Opt) Bool() bool {
	return o.Value.Bool()
}

func (o Opt) Int() int {
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

// OptName is the type of the option name group.
type OptName struct {
	Name      OptNameType
	ShortHand OptNameType
}

// Option Names definition
var (
	request    = OptName{Name: "request", ShortHand: "X"}
	form       = OptName{Name: "form", ShortHand: "F"}
	data       = OptName{Name: "data", ShortHand: "d"}
	fail       = OptName{Name: "fail", ShortHand: "f"}
	help       = OptName{Name: "help", ShortHand: "h"}
	include    = OptName{Name: "include", ShortHand: "i"}
	output     = OptName{Name: "output", ShortHand: "o"}
	remoteName = OptName{Name: "remote-name", ShortHand: "O"}
	silent     = OptName{Name: "silent", ShortHand: "s"}
	uploadFile = OptName{Name: "upload-file", ShortHand: "T"}
	user       = OptName{Name: "user", ShortHand: "u"}
	userAgent  = OptName{Name: "user-agent", ShortHand: "A"}
	verbose    = OptName{Name: "verbose", ShortHand: "v"}
	version    = OptName{Name: "version", ShortHand: "V"}
)

// options is the slice of the options.
var options = []Opt{
	{
		Name:      request.Name,
		ShortHand: request.ShortHand,
		Help:      "Specify request method to use",
		Usage:     "--request <method>",
		Required:  false,
		Supported: true,
		Value:     *NewOptValue(String),
	},
	{
		Name:      form.Name,
		ShortHand: form.ShortHand,
		Help:      "Specify multipart MIME data",
		Usage:     "--form <name=content>",
		Required:  false,
		Supported: true,
		Value:     *NewOptValue(String),
	},
	{
		Name:      data.Name,
		ShortHand: data.ShortHand,
		Help:      "HTTP POST data",
		Usage:     "--data <data>",
		Required:  false,
		Supported: true,
		Value:     *NewOptValue(String),
	},
	{
		Name:      fail.Name,
		ShortHand: fail.ShortHand,
		Help:      "Fail silently (no output at all) on HTTP errors",
		Usage:     "--fail",
		Required:  false,
		Supported: true,
		Value:     *NewOptValue(Bool),
	},
	{
		Name:      help.Name,
		ShortHand: help.ShortHand,
		Help:      "Show this help message and exit",
		Usage:     "--help <category>",
		Required:  false,
		Supported: true,
		Value:     *NewOptValue(Bool),
	},
	{
		Name:      include.Name,
		ShortHand: include.ShortHand,
		Help:      "Include protocol headers in the output (H/F)",
		Usage:     "--include",
		Required:  false,
		Supported: true,
		Value:     *NewOptValue(Bool),
	},
	{
		Name:      output.Name,
		ShortHand: output.ShortHand,
		Help:      "Write to <file> instead of stdout",
		Usage:     "--output <file>",
		Required:  false,
		Supported: true,
		Value:     *NewOptValue(String),
	},
	{
		Name:      remoteName.Name,
		ShortHand: remoteName.ShortHand,
		Help:      "Write output to a file named as the remote file",
		Usage:     "--remote-name",
		Required:  false,
		Supported: false,
		Value:     *NewOptValue(Bool),
	},
	{
		Name:      silent.Name,
		ShortHand: silent.ShortHand,
		Help:      "Silent mode (don't output anything)",
		Usage:     "--silent",
		Required:  false,
		Supported: true,
		Value:     *NewOptValue(Bool),
	},
	{
		Name:      uploadFile.Name,
		ShortHand: uploadFile.ShortHand,
		Help:      "Transfer local FILE to destination",
		Usage:     "--upload-file <file>",
		Required:  false,
		Supported: false,
		Value:     *NewOptValue(String),
	},
	{
		Name:      user.Name,
		ShortHand: user.ShortHand,
		Help:      "Server user and password",
		Usage:     "--user <user:password>",
		Required:  false,
		Supported: true,
		Value:     *NewOptValue(String),
	},
	{
		Name:      userAgent.Name,
		ShortHand: userAgent.ShortHand,
		Help:      "User-Agent to send to server",
		Usage:     "--user-agent <name>",
		Required:  false,
		Supported: true,
		Value:     *NewOptValue(String),
	},
	{
		Name:      verbose.Name,
		ShortHand: verbose.ShortHand,
		Help:      "Make the operation more talkative",
		Usage:     "--verbose",
		Required:  false,
		Supported: true,
		Value:     *NewOptValue(Bool),
	},
	{
		Name:      version.Name,
		ShortHand: version.ShortHand,
		Help:      "Show version number and quit",
		Usage:     "--version",
		Required:  false,
		Supported: true,
		Value:     *NewOptValue(Bool),
	},
}

// Init initializes the parser.
func ParserInit() {
	// Bind the optional arguments to the options (variables).
	for i, opt := range options {
		// NOTE: We use the index of the option in the slice as the flag name.
		// If we don't access the flag by index, the flag value will not be updated.
		// This is Go language's behavior.
		// In order to update array elements, we need to use the index.
		if opt.Value.GetType() == String {
			flag.StringVar(&options[i].Value.strValue, opt.Name.String(), "", opt.GetLineToPrint())
			flag.StringVar(&options[i].Value.strValue, opt.ShortHand.String(), "", opt.GetLineToPrint())
		} else if opt.Value.GetType() == Bool {
			flag.BoolVar(&options[i].Value.boolValue, opt.Name.String(), false, opt.GetLineToPrint())
			flag.BoolVar(&options[i].Value.boolValue, opt.ShortHand.String(), false, opt.GetLineToPrint())
		} else if opt.Value.GetType() == Int {
			flag.IntVar(&options[i].Value.intValue, opt.Name.String(), 0, opt.GetLineToPrint())
			flag.IntVar(&options[i].Value.intValue, opt.ShortHand.String(), 0, opt.GetLineToPrint())
		}
	}
	flag.Parse()
}

// GetOptSize returns the number of options.
func GetOptSize() int {
	return len(options)
}

// GetSupportedOptSize returns the number of supported options.
func GetSupportedOptSize() int {
	size := 0
	for _, opt := range options {
		if opt.Supported {
			size++
		}
	}
	return size
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
