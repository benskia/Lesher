package power

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// Description:
//	Manages information about available batteries (per Linux power_supply
//	class).
//
//	Active batteries and their files might change between executions, so we
//	probably don't need persisting information about the batteries.
//
//	Lesher sets the same thresholds for all batteries, but there might be cases
//	(such as before the first profile is set) when batteries have different
//	thresholds. Still considering if it makes sense to print details for every
//	battery individually.
//
// Responsibilities:
//	- Read current charge thresholds.
//	- Write new charge thresholds.
//	- Read battery health information.
//	- Calculate battery health.

// Reads power_supply info into Batteries that ops can work with.
func GetThresholds() ([]battery, error) {
	batteries, err := getPowerSupplies(batFilepath)
	if err != nil {
		return nil, fmt.Errorf("failed to find power supplies: %v", err)
	}

	for _, bat := range batteries {
		err := bat.readThresholds()
		if err != nil {
			return nil, fmt.Errorf("failed to read %s thresholds: %v", bat.name, err)
		}
	}

	return batteries, nil
}

// Writes power supply info for all power_supplies passed by ops.
func SaveThresholds(batteries []battery) error {
	for _, bat := range batteries {
		if err := bat.writeThresholds(); err != nil {
			return fmt.Errorf("failed to write %s thresholds: %v", bat.name, err)
		}
	}
	return nil
}

// Returns a slice of active batteries with names to populate with info.
func getPowerSupplies(filepath string) ([]battery, error) {
	dirs, err := os.ReadDir(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed ReadDir: %v", err)
	}

	// ./power_supply/* should only contain BAT# and possibly AC. We only need BATs.
	batteries := []battery{}
	for _, dir := range dirs {
		if strings.Contains(dir.Name(), "BAT") {
			batteries = append(batteries, battery{name: dir.Name()})
		}
	}

	if len(batteries) == 0 {
		return nil, errors.New("no batteries found")
	}

	return batteries, nil
}
