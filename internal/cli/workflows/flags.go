package workflows

import (
	"strconv"
	"strings"

	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/log"
)

var GlobalFlagsToArgs = map[string]string{
	flag.Profile:      "1",
	flag.ProfileShort: "1",
	flag.Debug:        "0",
	flag.DebugShort:   "0",
}

func shouldRemoveFlagAndArgs(flags map[string]string, arg string) (int, error) {
	trimmedArg := strings.TrimLeft(arg, "-")

	flagArgsCount := flags[trimmedArg]
	if flagArgsCount == "" {
		flagArgsCount = GlobalFlagsToArgs[trimmedArg]
	}

	if flagArgsCount != "" {
		flagArgsCountValue, err := strconv.Atoi(flagArgsCount)
		if err != nil {
			return -1, err
		}
		return flagArgsCountValue, nil
	}

	return -1, nil

}

func RemoveFlagsAndArgs(flagsToBeRemoved map[string]string, argsToBeRemoved map[string]bool, args []string) ([]string, error) {
	var newArgs []string
	for i := 0; i < len(args); i++ {
		arg := args[i]

		// remove global flags
		flagArgsCount, err := shouldRemoveFlagAndArgs(flagsToBeRemoved, arg)
		if err != nil {
			_, _ = log.Warningf("Error while removing flags and arg(s): %s\n", err)
			return nil, err
		}

		if flagArgsCount >= 0 {
			_, _ = log.Warningf("Skipped flag %s and %d args\n", args[i], flagArgsCount)
			i = i + flagArgsCount // skip the args after flag
			continue
		}

		// remove args
		if argsToBeRemoved[arg] {
			_, _ = log.Debugf("Skipped arg %s\n", args[i])
			continue
		}

		newArgs = append(newArgs, arg)
	}
	return newArgs, nil
}
