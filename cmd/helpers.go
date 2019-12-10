package cmd

import (
	"encoding/json"
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

func configDir() string {
	home, err := configHome()
	if err != nil {
		exitOnErr(err)
	}
	return home
}

func exitOnErr(msg interface{}) {
	if msg != nil {
		fmt.Println("Error:", msg)
		os.Exit(1)
	}
}

func prettyJSON(obj interface{}) {
	prettyJSON, err := json.MarshalIndent(obj, "", "\t")
	exitOnErr(err)
	fmt.Println(string(prettyJSON))
}
