package main

import (
	"fmt"
	"os"

	"github.com/sota0121/gurl/cmd/feature"
)

func main() {
	command := feature.NewCommand()
	if command == nil {
		fmt.Println("command is nil")
		os.Exit(1)
	}
	if err := command.Execute(); err != nil {
		fmt.Println(err)
	}
}
