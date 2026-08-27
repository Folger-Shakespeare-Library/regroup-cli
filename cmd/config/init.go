package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/Folger-Shakespeare-Library/regroup-cli/internal/cfg"
	"golang.org/x/term"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Set up Regroup API credentials interactively",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := cfg.Load(); err != nil {
			return err
		}

		reader := bufio.NewReader(os.Stdin)

		apiKey, err := prompt(reader, "API Key", viper.GetString("api_key"))
		if err != nil {
			return err
		}
		viper.Set("api_key", apiKey)

		apiSecret, err := promptSecret("API Secret")
		if err != nil {
			return err
		}
		if apiSecret != "" {
			viper.Set("api_secret", apiSecret)
		}

		if err := cfg.Save(); err != nil {
			return err
		}
		fmt.Println("Configuration saved.")
		return nil
	},
}

func prompt(reader *bufio.Reader, label, current string) (string, error) {
	if current != "" {
		fmt.Printf("%s [%s]: ", label, current)
	} else {
		fmt.Printf("%s: ", label)
	}
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	input = strings.TrimSpace(input)
	if input == "" {
		return current, nil
	}
	return input, nil
}

func promptSecret(label string) (string, error) {
	fmt.Printf("%s (input hidden): ", label)
	b, err := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(b)), nil
}
