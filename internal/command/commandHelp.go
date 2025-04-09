package command

import (
	"fmt"

	"github.com/benskia/Thresher/internal/config"
)

const helpDescription string = `Usage: Thresher help
	Displays Thresher documentation.
`

func commandHelp(_ *config.Config, _ []string) error {
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
	fmt.Print(`
Users can run Thresher to list battery threshold stats, check fullCharge health,
create threshold profiles, and activate existing profiles. This is done using
charge_control files of the Linux power_supply class.

`)
	for _, cmd := range GetCommands() {
		fmt.Println(cmd.Name)
		fmt.Println(cmd.Description)
		fmt.Println()
	}

	return nil
}
