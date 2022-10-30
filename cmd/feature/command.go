package feature

import (
	"flag"
	"fmt"
	"log"
)

// GurlCommand is the main command.
type GurlCommand struct {
	ctx *ReqContext
	cfg *CmdConfig

	client string // dummy field
	usage  string // dummy field
	url    string
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
			str := opt.String()
			fmt.Println(prefix, str)
		case Bool:
			b := opt.Bool()
			fmt.Println(prefix, b)
		case Int:
			i := opt.Int()
			fmt.Println(prefix, i)
		}
		fmt.Println()
	}
	return nil
}

// NewCommand creates a new command.
func NewCommand() *GurlCommand {
	// Initialize Flag Command Line Arguments
	ParserInit()

	// Get Options for Request Context
	optRequest, err := GetOptWithName(request)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	optForm, err := GetOptWithName(form)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	optData, err := GetOptWithName(data)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	optUploadFile, err := GetOptWithName(uploadFile)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	optUser, err := GetOptWithName(user)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	optUserAgent, err := GetOptWithName(userAgent)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	// Get Options for Command Config
	optHelp, err := GetOptWithName(help)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	optFail, err := GetOptWithName(fail)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	optInclude, err := GetOptWithName(include)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	optOutput, err := GetOptWithName(output)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	optRemoteName, err := GetOptWithName(remoteName)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	optSilent, err := GetOptWithName(silent)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	optVerbose, err := GetOptWithName(verbose)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	optVersion, err := GetOptWithName(version)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	// Get Positional Arguments
	positionalArgs := flag.Args()
	if len(positionalArgs) == 0 {
		log.Fatal("no url specified")
		return nil
	}
	url := positionalArgs[0]

	return &GurlCommand{
		ctx: NewReqContext(
			optRequest.Value.String(),
			optForm.Value.String(),
			optData.Value.String(),
			optUploadFile.Value.String(),
			optUser.Value.String(),
			optUserAgent.Value.String(),
		),
		cfg: NewCmdConfig(
			optHelp.Value.Bool(),
			optFail.Value.Bool(),
			optInclude.Value.Bool(),
			optRemoteName.Value.Bool(),
			optSilent.Value.Bool(),
			optVerbose.Value.Bool(),
			optVersion.Value.Bool(),
			optOutput.Value.String(),
		),
		client: "dummy",
		usage:  "Usage: gurl [options...] <url>",
		url:    url,
	}
}
