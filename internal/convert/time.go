// Copyright 2023 MongoDB Inc
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

package convert

import (
	"fmt"
	"time"
)

var ErrInvalidTimestamp = fmt.Errorf("supported timestamp formats examples: %s, 2006-01-02T15:04:05-0700", time.RFC3339)

func ParseTimestamp(timestamp string) (time.Time, error) {
	layouts := []string{
		time.RFC3339,
		"2006-01-02T15:04:05-0700",
	}

	var parsedTime time.Time
	var err error

	for _, layout := range layouts {
		parsedTime, err = time.Parse(layout, timestamp)
		if err == nil {
			return parsedTime, err
		}
	}
	return parsedTime, fmt.Errorf("%w, %w", err, ErrInvalidTimestamp)
}
