package config

import (
	"errors"
	"fmt"
	"os"
)

func configHome() (string, error) {
	if home := os.Getenv("XDG_CONFIG_HOME"); home != "" {
		return home, nil
	}
	if home := os.Getenv("HOME"); home != "" {
		return fmt.Sprintf("%s/.config", home), nil
	}

	return "", errors.New("could not find home")
}
