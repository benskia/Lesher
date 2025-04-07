package command

import (
	"fmt"

	"github.com/benskia/Lesher/internal/config"
	"github.com/benskia/Lesher/internal/power"
)

const listDescription string = `Usage: Lesher list
	Lists existing profiles and active battery thresholds.
`

func commandList(cfg *config.Config, _ []string) error {
	fmt.Println()
	fmt.Println("Profiles:")
	for _, profile := range cfg.Profiles {
		fmt.Printf("\tName: %s\n", profile.Name)
		fmt.Printf("\tStart: %d\tEnd: %d\n\n", profile.Start, profile.End)
	}

	batteries, err := power.GetThresholds()
	if err != nil {
		return err
	}

	fmt.Println("Current Thresholds:")
	for _, battery := range batteries {
		fmt.Printf("\tName: %s\n", battery.Name)
		fmt.Printf("\tStart: %d\tEnd: %d\n\n", battery.Start, battery.End)
	}

	return nil
}
