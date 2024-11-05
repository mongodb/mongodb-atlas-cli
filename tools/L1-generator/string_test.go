package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCleanString(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{
			`Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			`Flag that indicates whether the response body should be in the prettyprint format.`,
		},
		{
			`**Note**: This resource cannot be used to add a user invited via the deprecated [Invite One MongoDB Cloud User to Join One Project](#tag/Projects/operation/createProjectInvitation) endpoint.`,
			`Note: This resource cannot be used to add a user invited via the deprecated Invite One MongoDB Cloud User to Join One Project endpoint.`,
		},
		{
			`The delimiter that separates **databases.[n].collections.[n].dataSources.[n].path** segments in the data store. MongoDB Cloud uses the delimiter to efficiently traverse S3 buckets with a hierarchical directory structure. You can specify any character supported by the S3 object keys as the delimiter. For example, you can specify an underscore (_) or a plus sign (+) or multiple characters, such as double underscores (__) as the delimiter. If omitted, defaults to ` + "`" + `/` + "`" + `.`,
			`The delimiter that separates databases.[n].collections.[n].dataSources.[n].path segments in the data store. MongoDB Cloud uses the delimiter to efficiently traverse S3 buckets with a hierarchical directory structure. You can specify any character supported by the S3 object keys as the delimiter. For example, you can specify an underscore (_) or a plus sign (+) or multiple characters, such as double underscores (__) as the delimiter. If omitted, defaults to /.`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			output, err := Clean(tt.input)

			require.NoError(t, err)
			require.Equal(t, tt.output, output)
		})
	}
}
