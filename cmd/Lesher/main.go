package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/benskia/Lesher/internal/config"
)

// Users can run Lesher to list battery threshold stats, check fullCharge
// health, create threshold profiles, and activate existing profiles. This is
// done using charge_control files of the Linux power_supply class.
//
// Usage: Lesher <cmd> [arg...]
//
// help
//		Display this documentation.
//
// list
//		Lists available profiles. Indicates active profile if any.
//
// health
//		For each battery, displays the possible full-charge as a percentage of
//		the battery's full-charge design specification.
//
// create <name> <start> <end>
//		Creates or overwrites profile <name> that starts charging at <start>
//		percent and stops at <end>.
//
// set <name>
//		Sets profile <name> as the active profile.

// TODO: powerSupply package, cmds

func main() {
	powerSupplyDir := "/sys/class/power_supply/"
	batteries, err := getDirs(powerSupplyDir)
	if err != nil {
		log.Fatalf("failed to get power supplies\n%v", err)
	}
	for _, battery := range batteries {
		fmt.Println(battery)
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cfg.Profiles)
}

func getDirs(filepath string) ([]string, error) {
	dirs, err := os.ReadDir(filepath)
	if err != nil {
		return nil, fmt.Errorf("getDirs: %v", err)
	}

	// ./power_supply/* should only contain BAT# and possibly AC. We only need BATs.
	batteries := []string{}
	for _, dir := range dirs {
		if strings.Contains(dir.Name(), "BAT") {
			batteries = append(batteries, dir.Name())
		}
	}

	if len(batteries) == 0 {
		return nil, errors.New("getDirs: no batteries found")
	}

	return batteries, nil
}
