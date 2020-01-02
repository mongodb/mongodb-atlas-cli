package config

import (
	"strings"
	"testing"

	"github.com/spf13/viper"
)

func TestProfile_GetAPIPath(t *testing.T) {
	p := New("test")
	viper.Set(opsManagerURL, "example")
	if !strings.Contains(p.GetAPIPath(), publicAPIPath) {
		t.Errorf("GetAPIPath() = %s; want '%s'", p.GetAPIPath(), publicAPIPath)
	}
	viper.Set(service, CloudService)
	if !strings.Contains(p.GetAPIPath(), atlasAPIPath) {
		t.Errorf("GetAPIPath() = %s; want '%s'", p.GetAPIPath(), atlasAPIPath)
	}
}
