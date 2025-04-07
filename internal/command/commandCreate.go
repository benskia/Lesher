package command

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/benskia/Lesher/internal/config"
)

const createDescription string = `Usage: Lesher create <name> <start> <end>
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

	// We can just update the existing values if the profile already exists.
	if profile, ok := cfg.Profiles[name]; ok {
		fmt.Printf("Profile %s found. Updating values ...\n", name)
		profile.Start = start
		profile.End = end
		cfg.Profiles[name] = profile

		if err := cfg.SaveConfig(); err != nil {
			return fmt.Errorf("error saving profiles: %v", err)
		}

		return nil
	}

	fmt.Printf("Creating profile %s with start %d and end %d ...\n", name, start, end)
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
