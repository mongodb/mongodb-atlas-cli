package config

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
)

func configHome() (string, error) {
	if home := os.Getenv("XDG_CONFIG_HOME"); home != "" {
		return home, nil
	}
	home, err := homedir.Dir()

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/.config", home), nil
}
