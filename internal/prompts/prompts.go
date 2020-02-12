package prompts

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

// NewDeleteConfirmation creates a prompt to confirm if the entry should be deleted
func NewDeleteConfirmation(entry string) *survey.Confirm {
	prompt := &survey.Confirm{
		Message: fmt.Sprintf("Are you sure you want to delete: %s", entry),
	}
	return prompt
}
