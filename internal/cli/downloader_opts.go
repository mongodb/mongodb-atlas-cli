// Copyright 2020 MongoDB Inc
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

package cli

import (
	"io"
	"os"

	"github.com/spf13/afero"
)

// DownloaderOpts options required when deleting a resource.
// A command can compose this struct and then safely rely on the methods Prompt, or Delete
// to manage the interactions with the user
type DownloaderOpts struct {
	Out string
	Fs  afero.Fs
}

func (opts *DownloaderOpts) NewWriteCloser() (io.WriteCloser, error) {
	// Create file only if is not there already (don't overwrite)
	ff := os.O_CREATE | os.O_TRUNC | os.O_WRONLY | os.O_EXCL
	f, err := opts.Fs.OpenFile(opts.Out, ff, 0777)
	return f, err
}

func (opts *DownloaderOpts) OnError(f io.Closer) error {
	_ = f.Close()
	return opts.Fs.Remove(opts.Out)
}
