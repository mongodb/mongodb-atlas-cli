package auth

import (
	"bytes"
	"testing"

	"github.com/mongodb/mongocli/internal/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWhoAmIBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		WhoAmIBuilder(),
		0,
		[]string{},
	)
}

func Test_whoOpts_Run(t *testing.T) {
	buf := new(bytes.Buffer)
	opts := &whoOpts{
		OutWriter: buf,
		account:   "test",
	}
	require.NoError(t, opts.Run())
	assert.Equal(t, "Logged in as test\n", buf.String())
}
