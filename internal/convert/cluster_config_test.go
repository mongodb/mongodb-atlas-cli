package convert

import "testing"

func TestClusterConfig_protocolVersion(t *testing.T) {
	c := ClusterConfig{Version: "4.0"}
	ver, err := c.protocolVer()

	if err != nil {
		t.Fatalf("protocolVer() unexpected error: %v", err)
	}
	if ver != "1" {
		t.Fatalf("protocolVer() expected: %s but got: %s", "1", ver)
	}

	c.Version = "3.2"
	ver, err = c.protocolVer()

	if err != nil {
		t.Fatalf("protocolVer() unexpected error: %v", err)
	}
	if ver != "0" {
		t.Fatalf("protocolVer() expected: %s but got: %s", "0", ver)
	}
}
