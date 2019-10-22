package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

func ValidateArgument(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return errors.New("One argument is required")
	}

	return nil
}
