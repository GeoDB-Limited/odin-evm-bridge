package main

import (
	"github.com/GeoDB-Limited/odin-evm-bridge/internal/cli"
	"os"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
