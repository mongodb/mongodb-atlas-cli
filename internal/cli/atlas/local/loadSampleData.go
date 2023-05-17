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

	"github.com/briandowns/spinner"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/spf13/cobra"
)

//go:embed files/dump.tar.gz
var dump []byte

type SampleDataOpts struct {
	cli.OutputOpts
	cli.GlobalOpts
	s *spinner.Spinner
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
		target := filepath.Join(output, header.Name)
		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(target, 0755); err != nil {
				return err
			}
		case tar.TypeReg:
			f, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			if err != nil {
				return err
			}
			defer f.Close()
			if _, err := io.Copy(f, r); err != nil {
				return err
			}
		}
	}
}

func (opts *SampleDataOpts) Run(ctx context.Context) error {
	if opts.s != nil {
		opts.s.Start()
	}

	defer func() {
		if opts.s != nil {
			opts.s.Stop()
		}
	}()

	dumpDir, err := os.MkdirTemp("", "dump")
	if err != nil {
		return err
	}
	defer os.RemoveAll(dumpDir)

	err = extractTarGz(dump, dumpDir)
	if err != nil {
		return err
	}

	cmd := exec.Command("mongorestore", "--gzip", "-u", localUser, "-p", localPassword, localUri, dumpDir)
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr
	// cmd.Stdin = os.Stdin
	err = cmd.Run()
	if err != nil {
		return err
	}

	if opts.s != nil {
		opts.s.Stop()
	}

	return opts.Print(nil)
}

// atlas local loadSampleData.
func SampleDataBuilder() *cobra.Command {
	opts := &SampleDataOpts{}
	if opts.IsTerminal() {
		opts.s = spinner.New(spinner.CharSets[9], speed)
	}
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

	return cmd
}
