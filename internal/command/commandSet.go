package command

import "github.com/benskia/Lesher/internal/config"

const setDescription string = `
Usage: Lesher set <name>
Activates profile <name> by writing its values to power_supply files.
`

func commandSet(cfg *config.Config, args []string) error {
	return nil
}
