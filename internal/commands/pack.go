package commands

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/plexsystems/pacmod/pack"
	"github.com/spf13/cobra"
)

// NewPackCommand creates a new pack command which allows
// the user to package their Go modules
func NewPackCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:   "pack <module>",
		Short: "Package your Go modules",
		Args:  cobra.MinimumNArgs(1),

		RunE: func(cmd *cobra.Command, args []string) error {
			return runPackCommand(args)
		},
	}

	return &cmd
}

func runPackCommand(args []string) error {
	path, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("get working directory: %w", err)
	}

	module := args[0]

	path, err = filepath.Abs(path)
	if err != nil {
		return fmt.Errorf("get abs path of module path: %w", err)
	}

	log.Printf("Packing %s...", module)
	if err := pack.Module(module); err != nil {
		return fmt.Errorf("package module: %w", err)
	}

	return nil
}
