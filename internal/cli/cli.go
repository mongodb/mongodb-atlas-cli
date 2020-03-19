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
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/prompts"
)

const (
	fallbackSuccessMessage = "'%s' deleted\n"
	fallbackFailMessage    = "'%s' not deleted\n"
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
// A command can embed this structure and then safely rely on the methods Confirm, DeleteFromProject or Delete
// to manage the interactions with the user
type deleteOpts struct {
	entry          string
	confirm        bool
	successMessage string
	failMessage    string
}

// DeleterFromProject a function to delete from the store.
type DeleterFromProject func(projectID string, entry string) error

// DeleterFromProjectAuthDB a function to delete from the store.
type DeleterFromProjectAuthDB func(authDB string, projectID string, entry string) error

// DeleteFromProject deletes a resource from a project, it expects a callback
// that should perform the deletion from the store.
func (opts *deleteOpts) DeleteFromProject(d DeleterFromProject, projectID string) error {
	if !opts.confirm {
		opts.printFailMessage()
		return nil
	}
	err := d(projectID, opts.entry)

	if err != nil {
		return err
	}

	opts.printSuccessMessage()

	return nil
}

// DeleterFromProjectAuthDB deletes a resource from a project, it expects a callback
// that should perform the deletion from the store.
func (opts *deleteOpts) DeleterFromProjectAuthDB(d DeleterFromProjectAuthDB, authDB, projectID string) error {
	if !opts.confirm {
		opts.printFailMessage()
		return nil
	}
	err := d(authDB, projectID, opts.entry)

	if err != nil {
		return err
	}

	opts.printSuccessMessage()

	return nil
}

// printSuccessMessage prints a success message
func (opts *deleteOpts) printSuccessMessage() {
	if opts.successMessage != "" {
		fmt.Printf(opts.successMessage, opts.entry)
	} else {
		fmt.Printf(fallbackSuccessMessage, opts.entry)
	}
}

// printFailMessage prints a fail message
func (opts *deleteOpts) printFailMessage() {
	if opts.successMessage != "" {
		fmt.Printf(opts.failMessage, opts.entry)
	} else {
		fmt.Printf(fallbackFailMessage, opts.entry)
	}
}

// Deleter a function to delete from the store.
type Deleter func(entry string) error

// Delete deletes a resource not associated to a project, it expects a callback
//// that should perform the deletion from the store.
func (opts *deleteOpts) Delete(d Deleter) error {
	if !opts.confirm {
		opts.printFailMessage()
		return nil
	}
	err := d(opts.entry)

	if err != nil {
		return err
	}

	opts.printSuccessMessage()

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
