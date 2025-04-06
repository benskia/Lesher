package main

import (
	"fmt"
	"os"

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
// delete <name>
//		Deletes profile <name> if it exists.
//
// set <name>
//		Sets profile <name> as the active profile.

func main() {
	// Load/Create config
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("%v\nCreating new config...\n", err)
		if err := cfg.SaveConfig(); err != nil {
			fmt.Println(err)
		}
	}

	// Execute ops
	if len(os.Args) == 0 {
		fmt.Println("Missing args. Try: Lesher help")
	}

	fmt.Println(cfg.Profiles)
}
