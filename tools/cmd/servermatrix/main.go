// Copyright 2026 MongoDB Inc
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

// servermatrix updates build/ci/server-matrix.json from the Server release feed.
package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/tools/cmd/genevergreen/generate"
)

func run(ctx context.Context, w io.Writer) error {
	var feedURL, matrixPath, exceptionsPath string
	flag.StringVar(&feedURL, "feed", "https://downloads.mongodb.org/full.json", "Server release feed URL")
	flag.StringVar(&matrixPath, "matrix", "build/ci/server-matrix.json", "server matrix to update")
	flag.StringVar(&exceptionsPath, "exceptions", "build/ci/server-matrix-exceptions.yaml", "pairs where Atlas CLI differs from Server")
	flag.Parse()

	current, err := generate.LoadServerMatrix(matrixPath)
	if err != nil {
		return err
	}
	b, err := os.ReadFile(exceptionsPath)
	if err != nil {
		return err
	}
	exceptions, err := parseExceptions(b)
	if err != nil {
		return fmt.Errorf("%s: %w", exceptionsPath, err)
	}
	f, err := fetchFeed(ctx, feedURL)
	if err != nil {
		return err
	}
	res, err := update(current, f, exceptions)
	if err != nil {
		return err
	}
	out, err := marshalMatrix(res.Matrix)
	if err != nil {
		return err
	}
	const permissions = 0o600
	if err := os.WriteFile(matrixPath, out, permissions); err != nil {
		return err
	}
	printSummary(w, res)
	return nil
}

func printSummary(w io.Writer, res *result) {
	if len(res.Added) == 0 && len(res.Removed) == 0 {
		fmt.Fprintln(w, "server matrix unchanged")
	}
	for _, sv := range res.AddedVersions {
		fmt.Fprintf(w, "added version %s: %s\n", sv, strings.Join(res.Matrix[sv], ", "))
	}
	for _, p := range res.Added {
		fmt.Fprintf(w, "added %s\n", p)
	}
	for _, p := range res.Removed {
		fmt.Fprintf(w, "removed %s\n", p)
	}
	for _, p := range res.Unknown {
		fmt.Fprintf(w, "unknown target %s, not added\n", p)
	}
}

func main() {
	if err := run(context.Background(), os.Stdout); err != nil {
		log.Fatal(err)
	}
}
