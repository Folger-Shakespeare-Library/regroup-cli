package config

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"regroup/internal/cfg"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Display current configuration",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := cfg.Load(); err != nil {
			return err
		}

		fmt.Printf("api_key:    %s\n", viper.GetString("api_key"))
		fmt.Printf("api_secret: %s\n", cfg.Mask(viper.GetString("api_secret")))
		return nil
	},
}
