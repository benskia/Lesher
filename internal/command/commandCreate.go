package command

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

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
		return fmt.Errorf("start must be a number: %w", err)
	}

	end, err := strconv.Atoi(args[2])
	if err != nil {
		return fmt.Errorf("end must be a number: %w", err)
	}

	if !(start < end) {
		return errors.New("start must be less than end")
	}

	// Updates are destructive, so we should get confirmation from the user.
	if _, ok := cfg.Profiles[name]; ok {
		fmt.Printf("Profile %s exists. Update? (y/N) ", name)
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		if err := scanner.Err(); err != nil {
			return fmt.Errorf("error confirming profile update: %w", err)
		}

		if strings.ToLower(scanner.Text()) != "y" {
			return errors.New("user cancelled update")
		}
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
