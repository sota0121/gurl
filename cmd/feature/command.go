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

	client *GurlClient
	usage  string
	url    string
}

func (c *GurlCommand) Execute() error {
	// Print Help
	if c.cfg.showHelp {
		fmt.Println(c.usage)
		return nil
	}

	// Print Option (for Debug)
	// for i := 0; i < GetOptSize(); i++ {
	// 	opt, err := GetOptWithIndex(i)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	fmt.Println(opt.GetLineToPrint())
	// 	prefix := "value: "
	// 	switch opt.GetType() {
	// 	case String:
	// 		str := opt.String()
	// 		fmt.Println(prefix, str)
	// 	case Bool:
	// 		b := opt.Bool()
	// 		fmt.Println(prefix, b)
	// 	case Int:
	// 		i := opt.Int()
	// 		fmt.Println(prefix, i)
	// 	}
	// 	fmt.Println()
	// }

	// Request
	resp, err := c.client.Request(c.ctx, c.url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Print Response
	fmt.Println(resp.Status)
	for k, v := range resp.Header {
		fmt.Println(k, v)
	}
	fmt.Println()
	fmt.Println(resp.Body)

	return nil
}

// NewCommand creates a new command.
func NewCommand() *GurlCommand {
	// Initialize Flag Command Line Arguments
	ParserInit()

	// Get Options for Request Context
	optRequest, err := GetOptWithName(request.Name)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	optForm, err := GetOptWithName(form.Name)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	optData, err := GetOptWithName(data.Name)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	optUploadFile, err := GetOptWithName(uploadFile.Name)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	optUser, err := GetOptWithName(user.Name)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	optUserAgent, err := GetOptWithName(userAgent.Name)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	// Get Options for Command Config
	optHelp, err := GetOptWithName(help.Name)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	optFail, err := GetOptWithName(fail.Name)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	optInclude, err := GetOptWithName(include.Name)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	optOutput, err := GetOptWithName(output.Name)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	optRemoteName, err := GetOptWithName(remoteName.Name)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	optSilent, err := GetOptWithName(silent.Name)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	optVerbose, err := GetOptWithName(verbose.Name)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	optVersion, err := GetOptWithName(version.Name)
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

	// Create Request Context
	ctx := NewReqContext(
		optRequest.Value.String(),
		optForm.Value.String(),
		optData.Value.String(),
		optUploadFile.Value.String(),
		optUser.Value.String(),
		optUserAgent.Value.String())
	if ctx == nil {
		log.Fatal("failed to create request context")
		return nil
	}

	// Create Command Config
	cfg := NewCmdConfig(
		optHelp.Value.Bool(),
		optFail.Value.Bool(),
		optInclude.Value.Bool(),
		optRemoteName.Value.Bool(),
		optSilent.Value.Bool(),
		optVerbose.Value.Bool(),
		optVersion.Value.Bool(),
		optOutput.Value.String())
	if cfg == nil {
		log.Fatal("failed to create command config")
		return nil
	}

	// Create Gurl Client
	client := NewGurlClient()
	if client == nil {
		log.Fatal("failed to create gurl client")
		return nil
	}

	return &GurlCommand{
		ctx:    ctx,
		cfg:    cfg,
		client: client,
		usage:  "Usage: gurl [options...] <url>",
		url:    url,
	}
}
