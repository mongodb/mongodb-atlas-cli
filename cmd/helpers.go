package cmd

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

func exitOnErr(msg interface{}) {
	if msg != nil {
		fmt.Println("Error:", msg)
		os.Exit(1)
	}
}
