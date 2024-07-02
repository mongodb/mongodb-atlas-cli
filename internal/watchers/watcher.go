// Copyright 2023 MongoDB Inc
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

package watchers

import (
	"errors"
	"fmt"
	"time"

	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

type InvalidStateError struct {
	State             *string
	ErrorCode         *string
	ExpectedState     *string
	ExpectedErrorCode *string
}

func (err *InvalidStateError) Error() string {
	const (
		stateTemplate     = "Invalid state reached: %s."
		expectedTemplate  = "Expected state: %s."
		errorCodeTemplate = "Invalid error reached: %s."
		expectedErrorCode = "Expected error code: %s."
	)

	var got string
	var expected string

	if err.State != nil {
		got = fmt.Sprintf(stateTemplate, *err.State)
	} else {
		got = fmt.Sprintf(errorCodeTemplate, *err.ErrorCode)
	}

	if err.State != nil {
		expected = fmt.Sprintf(expectedTemplate, *err.State)
	} else {
		expected = fmt.Sprintf(expectedErrorCode, *err.ErrorCode)
	}

	return fmt.Sprintf("%s %s", got, expected)
}

type StateDescriber interface {
	GetState() (string, error)
}

type Watcher struct {
	Timeout            time.Duration // TODO: Timeout support - CLOUDP-181597
	NonConstantBackoff bool
	stateTransition    StateTransition
	describer          StateDescriber
	hasStarted         bool
	defaultWait        time.Duration
}

const defaultWait = 4 * time.Second

func NewWatcher(stateTransition StateTransition, describer StateDescriber) *Watcher {
	return NewWatcherWithDefaultWait(stateTransition, describer, defaultWait)
}

func NewWatcherWithDefaultWait(stateTransition StateTransition, describer StateDescriber, defaultWait time.Duration) *Watcher {
	return &Watcher{
		stateTransition: stateTransition,
		describer:       describer,
		defaultWait:     defaultWait,
	}
}

func (watcher *Watcher) Watch() error {
	watcher.hasStarted = false
	if watcher.NonConstantBackoff {
		return watcher.fibonacciBackoff()
	}

	return watcher.constantBackoff()
}

func (watcher *Watcher) fibonacciBackoff() error {
	previousBackoff := 0 * time.Second
	currentBackoff := watcher.defaultWait

	for {
		done, err := watcher.IsDone()
		if err != nil || done {
			return err
		}
		time.Sleep(currentBackoff)
		currentBackoff += previousBackoff
		previousBackoff = currentBackoff - previousBackoff
	}
}

func (watcher *Watcher) constantBackoff() error {
	for {
		done, err := watcher.IsDone()
		if err != nil || done {
			return err
		}
		time.Sleep(watcher.defaultWait)
	}
}

func (watcher *Watcher) IsDone() (bool, error) {
	if !watcher.stateTransition.HasStartState() {
		watcher.hasStarted = true
	}

	state, err := watcher.describer.GetState()

	if !watcher.hasStarted {
		if !watcher.stateTransition.IsStartState(state) {
			return false, &InvalidStateError{State: &state, ExpectedState: watcher.stateTransition.StartState}
		}
		watcher.hasStarted = true

		return false, nil
	}

	if err != nil {
		if watcher.stateTransition.IsRetryableError(err) {
			return false, nil
		} else if watcher.stateTransition.IsEndError(err) {
			return true, nil
		}
		return false, err
	}

	if watcher.stateTransition.IsEndState(state) {
		return true, nil
	} else if watcher.stateTransition.IsRetryableState(state) ||
		watcher.stateTransition.IsStartState(state) {
		return false, nil
	}

	if watcher.stateTransition.HasEndState() {
		return false, &InvalidStateError{State: &state, ExpectedState: watcher.stateTransition.EndState}
	}
	return false, &InvalidStateError{State: &state, ExpectedErrorCode: watcher.stateTransition.EndErrorCode}
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

func (st *StateTransition) HasEndState() bool {
	return st.EndState != nil
}

func (st *StateTransition) IsEndState(state string) bool {
	return st.EndState != nil && state == *st.EndState
}

func (st *StateTransition) IsEndError(err error) bool {
	if st.EndErrorCode == nil {
		return false
	}

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

func (st *StateTransition) HasEndError() bool {
	return st.EndErrorCode != nil
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
