package main

import (
	"fmt"

	"github.com/sota0121/gurl/cmd/feature"
)

func main() {
	command := feature.NewCommand()
	if err := command.Execute(); err != nil {
		fmt.Println(err)
	}
}
