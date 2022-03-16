package decryption

import (
	"bytes"
	"testing"
)

func Test_zeroLEK(t *testing.T) {
	d := DecryptSection{
		lek: []byte{1, 2, 3},
	}
	d.zeroLEK()
	if expected := []byte{0, 0, 0}; !bytes.Equal(d.lek, expected) {
		t.Errorf("expected: %v got: %v", expected, d.lek)
	}
}
