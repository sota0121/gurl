package feature

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/sota0121/gurl/internal"
)

// GurlCommand is the main command.
type GurlCommand struct {
	ctx *ReqContext
	cfg *CmdConfig

	client  *GurlClient
	usage   string
	version string
	url     string
}

// Execute executes the command.
func (c *GurlCommand) Execute() error {
	// ========================================
	// Early Exit
	// ========================================
	// -> Help
	if c.cfg.showHelp {
		fmt.Println(c.usage)
		return nil
	}
	// -> Version
	if c.cfg.version {
		fmt.Println(c.version)
		return nil
	}

	// ========================================
	// Initialize Command
	// ========================================
	level := Invalid
	if c.cfg.verbose {
		level = Verbose
	} else if c.cfg.silent {
		level = Silent
	} else {
		level = Normal
	}
	logger := NewLeveledLogger(level)
	logger.Println(c.cfg.verbose, "Logger Initialized with Level: ", level)

	// ========================================
	// Request Execution
	// ========================================
	// Request
	resp, err := c.client.Request(c.ctx, c.url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Handle Response
	// ... TODO

	// ========================================
	// Export Result
	// ========================================
	results := []string{}
	results = append(results, fmt.Sprintf("response Status: %s", resp.Status))

	if c.cfg.verbose {
		// Header
		for k, v := range resp.Header {
			results = append(results, fmt.Sprintf("%s: %s", k, v))
		}

		// Body
		results = append(results, fmt.Sprintf("response Body: %s", resp.Body))
	}

	// Output to file if specified
	if c.cfg.output != "" {
		f, err := os.Create(c.cfg.output)
		if err != nil {
			logger.Println(c.cfg.verbose, "Failed to create file: ", c.cfg.output)
			return err
		}
		defer f.Close()

		for _, result := range results {
			line := fmt.Sprintf("%s\n", result)
			_, err = f.WriteString(line)
			if err != nil {
				logger.Println(c.cfg.verbose, "Failed to write file: ", c.cfg.output)
				return err
			}
		}
	} else {
		for _, result := range results {
			logger.Println(c.cfg.verbose, result)
		}
	}
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

	// Get Positional Arguments
	url := ""
	// -> Are Positional Arguments Needed?
	needURL := !optHelp.Value.Bool() && !optVersion.Value.Bool()
	// -> Get Positional Arguments If Needed
	if needURL {
		positionalArgs := flag.Args()
		if len(positionalArgs) == 0 {
			log.Fatal("missing url")
			return nil
		}
		url = positionalArgs[0]
	}

	return &GurlCommand{
		ctx:     ctx,
		cfg:     cfg,
		client:  client,
		usage:   "Usage: gurl [options...] <url>",
		version: internal.Version,
		url:     url,
	}
}
