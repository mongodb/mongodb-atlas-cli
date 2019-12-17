package config

import (
	"os"
	"testing"
)

func TestConfig_configHome(t *testing.T) {
	_ = os.Setenv("XDG_CONFIG_HOME", "my_config")
	_ = os.Setenv("HOME", ".")

	home, err := configHome()
	if home != "my_config" {
		t.Errorf("configHome() = %s; want 'my_config'", home)
	}
	if err != nil {
		t.Fatalf("configHome() unexpected error: %v", err)
	}

	_ = os.Unsetenv("XDG_CONFIG_HOME")

	home, err = configHome()
	if home != "./.config" {
		t.Errorf("configHome() = %s; want './.config'", home)
	}
	if err != nil {
		t.Fatalf("configHome() unexpected error: %v", err)
	}
}
