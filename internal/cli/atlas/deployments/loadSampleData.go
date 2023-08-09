// Copyright 2023 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package local

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"context"
	_ "embed"
	"io"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/cobra"
)

//go:embed files/dump.tar.gz
var dump []byte

type SampleDataOpts struct {
	cli.OutputOpts
	cli.GlobalOpts
	debug bool
}

var sampleDataTemplate = `sample data loaded
`

func extractTarGz(input []byte, output string) error {
	gr, err := gzip.NewReader(bytes.NewReader(input))
	if err != nil {
		return err
	}
	defer gr.Close()

	r := tar.NewReader(gr)
	for {
		header, err := r.Next()
		switch {
		case err == io.EOF:
			return nil
		case err != nil:
			return err
		case header == nil:
			continue
		}
		target := filepath.Join(output, filepath.Clean(header.Name))
		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(target, os.ModePerm); err != nil {
				return err
			}
		case tar.TypeReg:
			f, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			if err != nil {
				return err
			}
			defer f.Close()
			if _, err := io.Copy(f, r); err != nil { //nolint:gosec // tar.gz is embedded inside binary, no way this will be tampered
				return err
			}
		}
	}
}

func (opts *SampleDataOpts) Run(_ context.Context) error {
	dumpDir, err := os.MkdirTemp("", "dump")
	if err != nil {
		return err
	}
	defer os.RemoveAll(dumpDir)

	err = extractTarGz(dump, dumpDir)
	if err != nil {
		return err
	}

	cmd := exec.Command("mongorestore", "--drop", "--gzip", "-u", localUser, "-p", localPassword, localURI, dumpDir)
	if opts.debug {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
	}
	err = cmd.Run()
	if err != nil {
		return err
	}

	return opts.Print(nil)
}

// atlas local loadSampleData.
func SampleDataBuilder() *cobra.Command {
	opts := &SampleDataOpts{}
	cmd := &cobra.Command{
		Use:   "loadSampleData",
		Short: "Loads sample data in a local instance.",
		Args:  require.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.InitOutput(cmd.OutOrStdout(), sampleDataTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}

	cmd.Flags().BoolVarP(&opts.debug, flag.Debug, flag.DebugShort, false, usage.Debug)

	return cmd
}
