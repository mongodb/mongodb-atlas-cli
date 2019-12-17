package config

import (
	"errors"
	"fmt"
	"log"
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

func createConfigFile() error {
	_, err := os.OpenFile(configFile(), os.O_RDONLY|os.O_CREATE, 0600)

	return err
}

func configFile() string {
	home, err := configHome()
	if err != nil {
		log.Fatalf("%v", err)
	}

	// TODO: We can now read but not write, see https://github.com/spf13/viper/pull/813
	return fmt.Sprintf("%s/%s.toml", home, Name)
}
