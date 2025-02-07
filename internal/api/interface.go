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

package api

import (
	"context"
	"io"
	"net/http"
)

type CommandExecutor interface {
	ExecuteCommand(ctx context.Context, commandRequest CommandRequest) (*CommandResponse, error)
}

type AccessTokenProvider interface {
	GetAccessToken() (string, error)
}

type ConfigProvider interface {
	GetBaseURL() (string, error)
}

type CommandConverter interface {
	ConvertToHTTPRequest(request CommandRequest) (*http.Request, error)
}

type CommandRequest struct {
	Command     Command
	Content     io.Reader
	ContentType string
	Format      string
	Parameters  map[string][]string
	Version     string
}

type CommandResponse struct {
	IsSuccess bool
	HTTPCode  int
	Output    io.ReadCloser
}

type ResponseFormatter interface {
	ContentType(format string) (string, error)
	Format(format string, readerCloser io.ReadCloser) (io.ReadCloser, error)
}
