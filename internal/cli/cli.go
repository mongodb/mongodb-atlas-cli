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
	"errors"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/prompts"
)

const (
	fallbackSuccessMessage = "'%s' deleted\n"
	fallbackFailMessage    = "entry not deleted"
)

type globalOpts struct {
	orgID     string
	projectID string
}

// newGlobalOpts returns an globalOpts
func newGlobalOpts() *globalOpts {
	return new(globalOpts)
}

// ProjectID returns the project id.
// If the id is empty, it caches it after querying config.
func (opts *globalOpts) ProjectID() string {
	if opts.projectID != "" {
		return opts.projectID
	}
	opts.projectID = config.ProjectID()
	return opts.projectID
}

// OrgID returns the organization id.
// If the id is empty, it caches it after querying config.
func (opts *globalOpts) OrgID() string {
	if opts.orgID != "" {
		return opts.orgID
	}
	opts.orgID = config.OrgID()
	return opts.orgID
}

// deleteOpts options required when deleting a resource.
// A command can compose this struct and then safely rely on the methods Confirm, or Delete
// to manage the interactions with the user
type deleteOpts struct {
	entry          string
	confirm        bool
	successMessage string
	failMessage    string
}

// Delete deletes a resource not associated to a project, it expects a callback
// that should perform the deletion from the store.
func (opts *deleteOpts) Delete(d interface{}, a ...string) error {
	if !opts.confirm {
		fmt.Println(opts.FailMessage())
		return nil
	}

	var err error
	fmt.Printf("%#v", d)
	switch f := d.(type) {
	case func(string) error:
		err = f(opts.entry)
	case func(string, string) error:
		err = f(a[0], opts.entry)
	case func(string, string, string) error:
		err = f(a[0], a[1], opts.entry)
	default:
		return errors.New("invalid")
	}

	if err != nil {
		return err
	}

	fmt.Printf(opts.SuccessMessage(), opts.entry)

	return nil
}

// Confirm confirms that the resource should be deleted
func (opts *deleteOpts) Confirm() error {
	if opts.confirm {
		return nil
	}

	prompt := prompts.NewDeleteConfirm(opts.entry)
	return survey.AskOne(prompt, &opts.confirm)
}

// SuccessMessage gets the set success message or the default value
func (opts *deleteOpts) SuccessMessage() string {
	if opts.successMessage != "" {
		return opts.successMessage
	}
	return fallbackSuccessMessage
}

// FailMessage gets the set fail message or the default value
func (opts *deleteOpts) FailMessage() string {
	if opts.failMessage != "" {
		return opts.failMessage
	}
	return fallbackFailMessage
}
