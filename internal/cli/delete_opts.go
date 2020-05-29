package cli

import (
	"errors"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mongodb/mongocli/internal/prompt"
)

const (
	fallbackSuccessMessage = "'%s' deleted\n"
	fallbackFailMessage    = "Entry not deleted"
)

// DeleteOpts options required when deleting a resource.
// A command can compose this struct and then safely rely on the methods Prompt, or Delete
// to manage the interactions with the user
type DeleteOpts struct {
	Entry      string
	Confirm    bool
	successMsg string
	failMsg    string
}

func NewDeleteOpts(successMsg, failMsg string) *DeleteOpts {
	return &DeleteOpts{
		successMsg: successMsg,
		failMsg:    failMsg,
	}
}

// Delete deletes a resource not associated to a project, it expects a callback
// that should perform the deletion from the store.
func (opts *DeleteOpts) Delete(d interface{}, a ...string) error {
	if !opts.Confirm {
		fmt.Println(opts.FailMessage())
		return nil
	}

	var err error
	switch f := d.(type) {
	case func(string) error:
		err = f(opts.Entry)
	case func(string, string) error:
		err = f(a[0], opts.Entry)
	case func(string, string, string) error:
		err = f(a[0], a[1], opts.Entry)
	default:
		return errors.New("invalid")
	}

	if err != nil {
		return err
	}

	fmt.Printf(opts.SuccessMessage(), opts.Entry)

	return nil
}

// Prompt confirms that the resource should be deleted
func (opts *DeleteOpts) Prompt() error {
	if opts.Confirm {
		return nil
	}

	p := prompt.NewDeleteConfirm(opts.Entry)
	return survey.AskOne(p, &opts.Confirm)
}

// SuccessMessage gets the set success message or the default value
func (opts *DeleteOpts) SuccessMessage() string {
	if opts.successMsg != "" {
		return opts.successMsg
	}
	return fallbackSuccessMessage
}

// FailMessage gets the set fail message or the default value
func (opts *DeleteOpts) FailMessage() string {
	if opts.failMsg != "" {
		return opts.failMsg
	}
	return fallbackFailMessage
}
