// Copyright 2022 MongoDB Inc
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

package decryption

import (
	"errors"
	"fmt"

	"github.com/klauspost/compress/zstd"
)

type CompressionMode string

const (
	CompressionModeNone CompressionMode = "none"
	CompressionModeZstd CompressionMode = "zstd"
)

var ErrUnsupportedCompression = errors.New("unsupported compression mode")

func decompress(compressionMode CompressionMode, src []byte) ([]byte, error) {
	switch compressionMode {
	case CompressionModeNone:
		return src, nil
	case CompressionModeZstd:
		decoder, err := zstd.NewReader(nil)
		if err != nil {
			return nil, err
		}
		return decoder.DecodeAll(src, nil)
	default:
		return nil, fmt.Errorf(`%w: %s`, ErrUnsupportedCompression, compressionMode)
	}
}
