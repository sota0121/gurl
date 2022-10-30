package cmd

// CmdConfig is the configuration of the command.
type CmdConfig struct {
	// showHelp is the flag corresponding to -h or --help.
	showHelp bool
	// silentFail is the flag corresponding to -f or --fail.
	silentFail bool
	// include is the flag corresponding to -i or --include.
	include bool
	// output is the flag corresponding to -o or --output.
	output string
	// remoteName is the flag corresponding to -O or --remote-name.
	remoteName bool
	// silent is the flag corresponding to -s or --silent.
	silent bool
	// verbose is the flag corresponding to -v or --verbose.
	verbose bool
	// version is the flag corresponding to -V or --version.
	version bool
}

// NewCmdConfig returns a new instance of CmdConfig.
func NewCmdConfig(
	optHelp,
	optFail,
	optInclude,
	optRemoteName,
	optSilent,
	optVerbose,
	optVersion bool,
	optOutput string) *CmdConfig {
	return &CmdConfig{
		showHelp:   optHelp,
		silentFail: optFail,
		include:    optInclude,
		output:     optOutput,
		remoteName: optRemoteName,
		silent:     optSilent,
		verbose:    optVerbose,
		version:    optVersion,
	}
}
