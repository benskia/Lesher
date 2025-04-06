package command

import (
	"fmt"

	"github.com/benskia/Lesher/internal/config"
	"github.com/benskia/Lesher/internal/power"
)

func commandHealth(_ *config.Config, _ []string) error {
	batteries, err := power.GetFullCharges()
	if err != nil {
		return err
	}

	fmt.Println("Battery Health:")
	for _, bat := range batteries {
		health := float64(bat.FullChargeActual) / float64(bat.FullChargeDesign)
		fmt.Printf("Name: %s\n", bat.Name)
		fmt.Printf("Full-Charge Actual: %d\tDesign: %d\n", bat.FullChargeActual, bat.FullChargeDesign)
		fmt.Printf("Health: %.2f\n", health)
	}

	return nil
}
