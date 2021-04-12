package mongosh

import (
	"fmt"
	"testing"
)

func TestBin(t *testing.T) {
	want := mongoshBin
	if isWindows() {
		want = fmt.Sprintf("%s.exe", mongoshBin)
	}
	if got := Bin(); got != want {
		t.Errorf("Bin() = %s, want %s", got, want)
	}
}
