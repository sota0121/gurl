package main

import (
	"fmt"
)

func cmdName() string {
	return "gurl"
}

func main() {
	cmdName := cmdName()
	fmt.Println("--------------------------------")
	fmt.Println("This is ", cmdName, " command")
	fmt.Println("--------------------------------")
	fmt.Println(`Usage: gurl [options...] <url>
	-d, --data <data>          HTTP POST data
	-f, --fail                 Fail silently (no output at all) on HTTP errors
	-h, --help <category>      Get help for commands
	-i, --include              Include protocol response headers in the output
	-o, --output <file>        Write to file instead of stdout
	-O, --remote-name          Write output to a file named as the remote file
	-s, --silent               Silent mode
	-T, --upload-file <file>   Transfer local FILE to destination
	-u, --user <user:password> Server user and password
	-A, --user-agent <name>    Send User-Agent <name> to server
	-v, --verbose              Make the operation more talkative
	-V, --version              Show version number and quit`)
}
