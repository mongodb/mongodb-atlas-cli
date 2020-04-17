package cli

import "testing"

func TestGetHostNameAndPort(t *testing.T) {
	t.Run("valid parameter", func(t *testing.T) {
		host, port, err := getHostNameAndPort("test:2000")
		if err != nil {
			t.Fatalf("getHostNameAndPort unexpecteted err: %#v\n", err)
		}
		if host != "test" {
			t.Errorf("Expected '%s', got '%s'\n", "test", host)
		}
		if port != 2000 {
			t.Errorf("Expected '%d', got '%d'\n", 2000, port)
		}
	})
	t.Run("incomplete format", func(t *testing.T) {
		_, _, err := getHostNameAndPort("test")
		if err == nil {
			t.Fatal("getHostNameAndPort should return an error\n")
		}
	})
	t.Run("incomplete format", func(t *testing.T) {
		_, _, err := getHostNameAndPort(":test")
		if err == nil {
			t.Fatal("getHostNameAndPort should return an error\n")
		}
	})
}

func TestGlobalOpts_PreRunE(t *testing.T) {
	t.Run("empty project ID", func(t *testing.T) {
		o := &globalOpts{}
		if err := o.PreRunE(); err != errMissingProjectID {
			t.Errorf("Expected err: %#v, got: %#v\n", errMissingProjectID, err)
		}
	})
	t.Run("invalid project ID", func(t *testing.T) {
		o := &globalOpts{projectID: "1"}
		if err := o.PreRunE(); err == nil {
			t.Errorf("Expected an error\n")
		}
	})
	t.Run("valid project ID", func(t *testing.T) {
		o := &globalOpts{projectID: "5e98249d937cfc52efdc2a9f"}
		if err := o.PreRunE(); err != nil {
			t.Fatalf("PreRunE() unexpected error %v\n", err)
		}
	})
}
