package config

import (
	"fmt"

	"github.com/spf13/cobra"
	"regroup/internal/cfg"
)

var pathCmd = &cobra.Command{
	Use:   "path",
	Short: "Print the config file path",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cfg.Path())
	},
}
