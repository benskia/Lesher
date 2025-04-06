package command

import (
	"fmt"

	"github.com/benskia/Lesher/internal/config"
	"github.com/benskia/Lesher/internal/power"
)

func commandList(cfg *config.Config, _ []string) error {
	fmt.Println("Profiles:")
	for _, profile := range cfg.Profiles {
		fmt.Printf("Name: %s\n", profile.Name)
		fmt.Printf("Start: %d\tEnd: %d\n", profile.Start, profile.End)
	}

	batteries, err := power.GetThresholds()
	if err != nil {
		return err
	}

	fmt.Println("Current Thresholds:")
	for _, battery := range batteries {
		fmt.Printf("Name: %s\n", battery.Name)
		fmt.Printf("Start: %d\tEnd: %d\n", battery.Start, battery.End)
	}

	return nil
}
