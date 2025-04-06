package command

import (
	"errors"
	"fmt"

	"github.com/benskia/Lesher/internal/config"
	"github.com/benskia/Lesher/internal/power"
)

const setDescription string = `
Usage: Lesher set <name>
Activates profile <name> by writing its values to power_supply files.
`

func commandSet(cfg *config.Config, args []string) error {
	if len(args) < 1 {
		return errors.New("set expects one arg: <name>")
	}

	name := args[0]

	profile, ok := cfg.Profiles[name]
	if !ok {
		return fmt.Errorf("profile %s not found", name)
	}

	if err := power.SaveThresholds(profile); err != nil {
		return fmt.Errorf("error saving thresholds: %v", err)
	}

	return nil
}
