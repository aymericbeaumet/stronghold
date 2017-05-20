package main

import (
	"fmt"
	"os"

	"github.com/aymericbeaumet/stronghold/lib/cli"
)

func main() {
	if err := cli.RootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.Exit(0)
}
