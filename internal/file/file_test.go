package file_test

import (
	"fmt"
	"testing"

	"github.com/mongodb/mongocli/internal/file"
	"github.com/spf13/afero"
)

func TestLoad(t *testing.T) {
	t.Run("file does not exists", func(t *testing.T) {
		appFS := afero.NewMemMapFs()
		filename := "test.xml"
		err := file.Load(appFS, filename, nil)
		if err == nil || err.Error() != fmt.Sprintf("file not found: %s", filename) {
			t.Errorf("Load() unexpected error: %v", err)
		}
	})
	t.Run("file with no ext", func(t *testing.T) {
		appFS := afero.NewMemMapFs()
		filename := "test"
		_ = afero.WriteFile(appFS, filename, []byte(""), 0600)
		err := file.Load(appFS, filename, nil)
		if err == nil || err.Error() != fmt.Sprintf("filename: %s requires valid extension", filename) {
			t.Errorf("Load() unexpected error: %v", err)
		}
	})
	t.Run("file with invalid ext", func(t *testing.T) {
		appFS := afero.NewMemMapFs()
		filename := "test.xml"
		_ = afero.WriteFile(appFS, filename, []byte(""), 0600)
		err := file.Load(appFS, filename, nil)
		if err == nil || err.Error() != "unsupported file type: xml" {
			t.Errorf("Load() unexpected error: %v", err)
		}
	})
	t.Run("valid json file", func(t *testing.T) {
		appFS := afero.NewMemMapFs()
		filename := "test.json"
		_ = afero.WriteFile(appFS, filename, []byte("{}"), 0600)
		out := new(map[string]interface{})
		err := file.Load(appFS, filename, out)
		if err != nil {
			t.Fatalf("Load() unexpected error: %v", err)
		}
	})
	t.Run("valid yaml file", func(t *testing.T) {
		appFS := afero.NewMemMapFs()
		filename := "test.yaml"
		_ = afero.WriteFile(appFS, filename, []byte(""), 0600)
		out := new(map[string]interface{})
		err := file.Load(appFS, filename, out)
		if err != nil {
			t.Fatalf("Load() unexpected error: %v", err)
		}
	})
}
