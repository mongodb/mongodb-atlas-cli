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
	Command    Command
	Content    io.Reader
	Format     string
	Parameters map[string][]string
	Version    string
}

type CommandResponse struct {
	IsSuccess bool
	Output    io.ReadCloser
}
