package command

import (
	"fmt"

	"github.com/gorillazer/ginny-cli/ginny/util"
)

// CheckArgs should be used to ensure the right command line arguments are
// passed before executing an example.
func CheckArgs(args []string) error {
	if len(args) < 1 {
		msg := "Missing required parameters"
		util.Error(msg)
		return fmt.Errorf(msg)
	}
	return nil
}
