package main

import (
	"os"

	"github.com/Folger-Shakespeare-Library/regroup-cli/cmd/root"
)

func main() {
	if err := root.Cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
