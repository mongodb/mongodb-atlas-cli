package teams

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	mocks "github.com/mongodb/mongodb-atlas-cli/internal/mocks/atlas"
	"github.com/mongodb/mongodb-atlas-cli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/internal/test"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20231115005/admin"
)

func TestRenameBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		RenameBuilder(),
		0,
		[]string{flag.TeamID, flag.OrgID},
	)
}

func Test_renameOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockTeamRenamer(ctrl)

	expected := &atlasv2.TeamResponse{
		Name: pointer.Get("test"),
	}

	opts := &renameOpts{
		store: mockStore,
		name:  "test",
	}

	mockStore.
		EXPECT().
		RenameTeam(opts.OrgID, opts.teamID, opts.newTeam()).
		Return(expected, nil).
		Times(1)

	require.NoError(t, opts.Run())
}
