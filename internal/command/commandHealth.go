package command

import (
	"fmt"

	"github.com/benskia/Lesher/internal/config"
	"github.com/benskia/Lesher/internal/power"
)

const healthDescription string = `Usage: Lesher health
	Lists full-charge stats for active batteries, and displays the remaining
	percentage of full-charge that is possible due to wear.
`

func commandHealth(_ *config.Config, _ []string) error {
	batteries, err := power.GetFullCharges()
	if err != nil {
		return err
	}

	fmt.Println("\nBattery Health:")
	for _, bat := range batteries {
		health := float64(bat.FullChargeActual) / float64(bat.FullChargeDesign)
		fmt.Printf("\tName: %s\n", bat.Name)
		fmt.Printf("\tFull-Charge Actual: %d\tDesign: %d\n", bat.FullChargeActual, bat.FullChargeDesign)
		fmt.Printf("\tHealth: %.2f%%\n\n", health*100)
	}

	return nil
}
