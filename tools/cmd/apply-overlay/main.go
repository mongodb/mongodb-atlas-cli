// Copyright 2024 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	_ "embed"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"slices"

	"github.com/speakeasy-api/openapi-overlay/pkg/overlay"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func main() {
	var (
		specPath    string
		overlayGlob string
	)

	var rootCmd = &cobra.Command{
		Use:   "apply-overlay",
		Short: "CLI which applies overlays into OpenAPI spec",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return run(specPath, overlayGlob, cmd.OutOrStdout())
		},
	}

	rootCmd.Flags().StringVar(&specPath, "spec", "", "Path to spec file")
	rootCmd.Flags().StringVar(&overlayGlob, "overlay", "", "Glob to overlay")
	_ = rootCmd.MarkFlagFilename("spec")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run(specPath, overlayGlob string, w io.Writer) error {
	var spec io.Reader = os.Stdin
	if specPath != "" {
		specFile, err := os.OpenFile(specPath, os.O_RDONLY, os.ModePerm)
		if err != nil {
			return err
		}

		spec = specFile

		defer specFile.Close()
	}

	files, err := filepath.Glob(overlayGlob)
	if err != nil {
		return err
	}

	slices.Sort(files)

	return applyOverlays(spec, files, w)
}

func applyOverlay(spec *yaml.Node, overlayPath string) error {
	overlayFile, err := os.OpenFile(overlayPath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return err
	}

	defer overlayFile.Close()

	var o overlay.Overlay
	dec := yaml.NewDecoder(overlayFile)

	if err := dec.Decode(&o); err != nil {
		return err
	}

	if err := o.Validate(); err != nil {
		return err
	}

	return o.ApplyTo(spec)
}

func applyOverlays(r io.Reader, overlayFilePaths []string, w io.Writer) error {
	if len(overlayFilePaths) == 0 {
		_, err := io.Copy(w, r)
		return err
	}

	var spec yaml.Node
	err := yaml.NewDecoder(r).Decode(&spec)
	if err != nil {
		return err
	}

	for _, overlayFilePath := range overlayFilePaths {
		if err := applyOverlay(&spec, overlayFilePath); err != nil {
			return err
		}
	}

	buf, err := yaml.Marshal(&spec)
	if err != nil {
		return err
	}

	_, err = w.Write(buf)
	return err
}
