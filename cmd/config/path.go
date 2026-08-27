package config

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/Folger-Shakespeare-Library/regroup-cli/internal/cfg"
)

var pathCmd = &cobra.Command{
	Use:   "path",
	Short: "Print the config file path",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cfg.Path())
	},
}
