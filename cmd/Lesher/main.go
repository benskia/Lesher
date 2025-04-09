package main

import (
	"fmt"
	"log"
	"os"

	"github.com/benskia/Lesher/internal/command"
	"github.com/benskia/Lesher/internal/config"
)

// Users can run Lesher to list battery threshold stats, check fullCharge
// health, create threshold profiles, and activate existing profiles. This is
// done using charge_control files of the Linux power_supply class.
//
// Usage: Lesher <cmd> [opts...]
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

	// Execute command
	if len(os.Args) < 2 {
		log.Fatal("Missing args. Try: Lesher help")
	}

	cmdName := os.Args[1]
	var cmdArgs []string
	if len(os.Args) > 2 {
		cmdArgs = os.Args[2:]
	}

	cmd, ok := command.GetCommands()[cmdName]
	if !ok {
		log.Fatal("Invalid command name. Try: Lesher help")
	}

	err = cmd.Callback(cfg, cmdArgs)
	if err != nil {
		log.Fatalf("error executing %s: %v", cmd.Name, err)
	}
}
