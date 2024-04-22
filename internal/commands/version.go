package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

const Version = "0.7.0"

func versionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Show goup version",
		RunE: func(c *cobra.Command, args []string) error {
			_, err := fmt.Printf("goup version v%s\n", Version)
			return err
		},
	}
}
