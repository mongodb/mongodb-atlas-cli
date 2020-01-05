package cli

import (
	"fmt"

	"github.com/10gen/mcli/internal/flags"
)

var errMissingProjectID = fmt.Errorf(`required flag(s) "%s" not set`, flags.ProjectID)
