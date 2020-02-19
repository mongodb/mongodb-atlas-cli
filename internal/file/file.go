package file

import (
	"encoding/json"
	"fmt"
	"path/filepath"

	"github.com/mongodb/mcli/internal/search"
	"github.com/spf13/afero"
	"gopkg.in/yaml.v2"
)

var supportedExts = []string{"json", "yaml", "yml"}

// Load loads a given filename into the out interface.
// The file should be a valid json or yaml format
func Load(fs afero.Fs, filename string, out interface{}) error {
	if exists, err := afero.Exists(fs, filename); !exists || err != nil {
		return fmt.Errorf("file not found: %s", filename)
	}
	ext := filepath.Ext(filename)
	if len(ext) <= 1 {
		return fmt.Errorf("filename: %s requires valid extension", filename)
	}
	configType := ext[1:]
	if !search.StringInSlice(supportedExts, configType) {
		return fmt.Errorf("unsupported file type: %s", configType)
	}
	file, err := afero.ReadFile(fs, filename)
	if err != nil {
		return err
	}

	switch configType {
	case "yaml", "yml":
		if err := yaml.Unmarshal(file, out); err != nil {
			return err
		}
	case "json":
		if err := json.Unmarshal(file, out); err != nil {
			return err
		}
	}

	return nil
}
