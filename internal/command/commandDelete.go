package command

import (
	"errors"
	"fmt"

	"github.com/benskia/Lesher/internal/config"
)

const deleteDescription string = `
Usage: Lesher delete <name>
Deletes profile <name> if it exists.
`

func commandDelete(cfg *config.Config, args []string) error {
	if len(args) < 1 {
		return errors.New("delete expects one arg: <name>")
	}

	name := args[0]
	defer delete(cfg.Profiles, name)

	if _, ok := cfg.Profiles[name]; !ok {
		return fmt.Errorf("profile %s not found", name)
	}

	fmt.Printf("Deleting profile %s ...\n", name)
	return nil
}
