package config

import (
	"strings"
	"testing"

	"github.com/spf13/viper"
)

func TestProfile_GetAPIPath(t *testing.T) {
	p := New("test")
	viper.Set(opsManagerURL, "example")
	if !strings.Contains(p.APIPath(), publicAPIPath) {
		t.Errorf("APIPath() = %s; want '%s'", p.APIPath(), publicAPIPath)
	}
	viper.Set(service, CloudService)
	if !strings.Contains(p.APIPath(), atlasAPIPath) {
		t.Errorf("APIPath() = %s; want '%s'", p.APIPath(), atlasAPIPath)
	}
}
