package cfg

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func Dir() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".config", "regroup")
}

func Path() string {
	return filepath.Join(Dir(), "config.json")
}

func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(Dir())

	viper.SetDefault("api_key", "")
	viper.SetDefault("api_secret", "")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil
		}
		return fmt.Errorf("reading config: %w", err)
	}
	return nil
}

func Save() error {
	dir := Dir()
	if err := os.MkdirAll(dir, 0700); err != nil {
		return fmt.Errorf("creating config dir: %w", err)
	}

	path := Path()
	if err := viper.WriteConfigAs(path); err != nil {
		return fmt.Errorf("writing config: %w", err)
	}
	return os.Chmod(path, 0600)
}

func Mask(s string) string {
	if len(s) <= 4 {
		return "****"
	}
	return "****" + s[len(s)-4:]
}
