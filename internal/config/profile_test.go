package config

import (
	"strings"
	"testing"

	"github.com/spf13/afero"

	"github.com/spf13/viper"
)

func TestProfile_GetAPIPath(t *testing.T) {
	p := &Profile{
		Name:      "test",
		configDir: "home",
		fs:        afero.NewMemMapFs(),
	}
	viper.Set(opsManagerURL, "example")
	if !strings.Contains(p.APIPath(), publicAPIPath) {
		t.Errorf("APIPath() = %s; want '%s'", p.APIPath(), publicAPIPath)
	}
	viper.Set(service, CloudService)
	if !strings.Contains(p.APIPath(), atlasAPIPath) {
		t.Errorf("APIPath() = %s; want '%s'", p.APIPath(), atlasAPIPath)
	}
}
