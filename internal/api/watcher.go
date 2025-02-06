package api

import (
	"context"
	"errors"
	"fmt"
	"time"
)

var (
	ErrWatcherFailedToBuildRequest        = errors.New("failed to build watcher request")
	ErrWatcherFailedToWatch               = errors.New("failed to watch")
	ErrWatcherFailedToExecuteWatchRequest = errors.New("failed to execute watch request")
)

const (
	WatcherWatchInterval = 1 * time.Second
)

type Watcher struct {
	executor CommandExecutor
}

func (w *Watcher) Watch(ctx context.Context, props WatcherProperties) error {
	// TODO: add timeout support
	//ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	//defer cancel()
	request, err := buildRequest(props)
	if err != nil || request == nil {
		return errors.Join(ErrWatcherFailedToBuildRequest, err)
	}

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			if err := watchInner(ctx, w.executor, *request); err != nil {
				return errors.Join(ErrWatcherFailedToWatch, err)
			}

			time.Sleep(WatcherWatchInterval)
		}
	}
}

func buildRequest(props WatcherProperties) (*CommandRequest, error) {
	panic(fmt.Sprintf("TODO. UNUSED:%v", props))
}

// Actual watcher logic without handling
func watchInner(ctx context.Context, executor CommandExecutor, commandRequest CommandRequest) error {
	response, err := executor.ExecuteCommand(ctx, commandRequest)
	if err != nil {
		return errors.Join(ErrWatcherFailedToExecuteWatchRequest, err)
	}

	panic(fmt.Sprintf("TODO. UNUSED:%v", response))
}
