package watchers

import (
	"errors"
	"fmt"
	"time"

	atlasv2 "go.mongodb.org/atlas-sdk/admin"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

type InvalidStateError struct {
	State string
}

func (err *InvalidStateError) Error() string {
	return fmt.Sprintf("Invalid state reached: %s", err.State)
}

type StatusDescriber interface {
	GetStatus() (string, error)
}

type Watcher struct {
	Timeout            time.Duration
	StateTransition    StateTransition
	Describer          StatusDescriber
	ExponentialBackoff bool
	hasStarted         bool
}

const defaultWait = 4 * time.Second

func (watcher *Watcher) Watch() error {
	if watcher.ExponentialBackoff {
		return watcher.exponentialBackoff()
	}

	return watcher.linearBackoff()
}

func (watcher *Watcher) exponentialBackoff() error {
	backoff := defaultWait
	for {
		done, err := watcher.IsDone()
		if err != nil || done {
			return err
		}
		time.Sleep(backoff)
		backoff *= 2
	}
}

func (watcher *Watcher) linearBackoff() error {
	for {
		done, err := watcher.IsDone()
		if err != nil || done {
			return err
		}
		time.Sleep(defaultWait)
	}
}

func (watcher *Watcher) IsDone() (bool, error) {
	if !watcher.StateTransition.HasStartState() {
		watcher.hasStarted = true
	}

	state, err := watcher.Describer.GetStatus()

	if !watcher.hasStarted {
		if watcher.StateTransition.IsStartState(state) {
			watcher.hasStarted = true
		}

		return false, nil
	}

	if err != nil {
		if watcher.StateTransition.IsRetryableError(err) {
			return false, nil
		} else if watcher.StateTransition.IsEndError(err) {
			return true, nil
		}
		return false, err
	}

	if watcher.StateTransition.IsEndState(state) {
		return true, nil
	} else if watcher.StateTransition.IsRetryableState(state) {
		return false, nil
	}
	return false, &InvalidStateError{State: state}
}

type StateTransition struct {
	StartState          *string
	EndState            *string
	EndErrorCode        *string
	RetryableStates     []string
	RetryableErrorCodes []string
}

func (st *StateTransition) HasStartState() bool {
	return st.StartState != nil
}

func (st *StateTransition) IsStartState(state string) bool {
	return st.StartState != nil && state == *st.StartState
}

func (st *StateTransition) IsEndState(state string) bool {
	return st.EndState != nil && state == *st.EndState
}

func (st *StateTransition) IsEndError(err error) bool {
	var atlasErr *atlas.ErrorResponse
	var atlasv2Err *atlasv2.GenericOpenAPIError
	var errCode string

	if st.EndErrorCode == nil {
		return false
	}

	switch {
	case errors.As(err, &atlasErr):
		errCode = atlasErr.ErrorCode
	case errors.As(err, &atlasv2Err):
		errCode = *atlasv2Err.Model().ErrorCode
	default:
		return false
	}

	return errCode == *st.EndErrorCode
}

func (st *StateTransition) IsRetryableError(err error) bool {
	var atlasErr *atlas.ErrorResponse
	var atlasv2Err *atlasv2.GenericOpenAPIError
	var errCode string

	switch {
	case errors.As(err, &atlasErr):
		errCode = atlasErr.ErrorCode
	case errors.As(err, &atlasv2Err):
		errCode = *atlasv2Err.Model().ErrorCode
	default:
		return false
	}

	for _, retryableErrCode := range st.RetryableErrorCodes {
		if retryableErrCode == errCode {
			return true
		}
	}
	return false
}

func (st *StateTransition) HasRetryableError() bool {
	return len(st.RetryableErrorCodes) > 0
}

func (st *StateTransition) IsRetryableState(state string) bool {
	for _, skipableState := range st.RetryableStates {
		if skipableState == state {
			return true
		}
	}
	return false
}

func (st *StateTransition) IsInvalidState(state string) bool {
	return !st.IsRetryableState(state) && !st.IsStartState(state) && !st.IsEndState(state)
}

func (st *StateTransition) InInvalidError(err error) bool {
	return !st.IsRetryableError(err) && !st.IsEndError(err)
}
