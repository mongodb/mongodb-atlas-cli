package e2e

import (
	"os"
	"path/filepath"
)

func Bin() (string, error) {
	cliPath, err := filepath.Abs(os.Getenv("E2E_BINARY"))
	if err != nil {
		return "", err
	}

	if _, err := os.Stat(cliPath); err != nil {
		return "", err
	}
	return cliPath, nil
}
