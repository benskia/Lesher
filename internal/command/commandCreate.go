package command

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/benskia/Thresher/internal/config"
)

const createDescription string = `Usage: Thresher create <name> <start> <end>
	Creates or overwrites profile <name> with the given <start> and <end> values.
`

func commandCreate(cfg *config.Config, args []string) error {
	if len(args) < 3 {
		return errors.New("create expects three args: <name> <start> <end>")
	}

	name := args[0]

	start, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("error converting start value: %v", err)
	}

	end, err := strconv.Atoi(args[2])
	if err != nil {
		return fmt.Errorf("error converting end value: %v", err)
	}

	if !(start < end) {
		return errors.New("start must be less than end")
	}

	// We can just update the existing values if the profile already exists.
	if _, ok := cfg.Profiles[name]; ok {
		fmt.Printf("Updating profile %s ...\n", name)
	} else {
		fmt.Printf("Creating profile %s ...\n", name)
	}

	cfg.Profiles[name] = config.Profile{
		Name:  name,
		Start: start,
		End:   end,
	}

	if err := cfg.SaveConfig(); err != nil {
		return fmt.Errorf("error saving profiles: %v", err)
	}

	return nil
}
