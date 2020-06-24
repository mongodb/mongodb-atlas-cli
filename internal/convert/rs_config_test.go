package convert

import "testing"

func TestProtocolVersion(t *testing.T) {
	testCases := map[string]struct {
		config          *RSConfig
		protocolVersion string
	}{
		"empty fcv": {
			config:          &RSConfig{},
			protocolVersion: "",
		},
		"post 4.0": {
			config:          &RSConfig{FCVersion: "4.0"},
			protocolVersion: "1",
		},
		"pre 4.0": {
			config:          &RSConfig{FCVersion: "3.6"},
			protocolVersion: "0",
		},
	}
	for name, tc := range testCases {
		m := tc.config
		expected := tc.protocolVersion
		t.Run(name, func(t *testing.T) {
			ver, err := m.protocolVer()
			if err != nil {
				t.Fatalf("protocolVer() unexpected error: %v\n", err)
			}
			if ver != expected {
				t.Errorf("protocolVer() expected: %s but got: %s", expected, ver)
			}
		})
	}
}
