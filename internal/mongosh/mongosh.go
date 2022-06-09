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
package mongosh

import (
	"os"
	"path"
	"syscall"

	exec "golang.org/x/sys/execabs"
)

func Detect() bool {
	return BinPath() != ""
}

func BinPath() string {
	if p, err := exec.LookPath(mongoshBin); err == nil {
		return path.Join(p, mongoshBin)
	}

	return ""
}

func Run(username, password, mongoURI string) error {
	args := []string{"-u", username, "-p", password, mongoURI}
	env := os.Environ()
	return syscall.Exec(BinPath(), args, env) //nolint:gosec // false positive, this path won't be tampered
}
