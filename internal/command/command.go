package command

import "github.com/benskia/Lesher/internal/config"

// Description:
//	Manages commands (ops) to execute as <args> [opts].
//
// Responsibilities:
//	- 'help': Documentation
//	- 'list': Current threshold values and profiles
//	- 'health': Percentage full-charge remaining
//	- 'create': Create or overwrite profile
//	- 'delete': Delete profile
//	- 'set': Set active profile

type Command struct {
	Name        string
	Description string
	Callback    func(cfg *config.Config, args []string) error
}

func GetCommands() map[string]Command {
	return map[string]Command{
		"help": {
			Name:        "help",
			Description: "Display program documentation",
			Callback:    commandHelp,
		},
		"list": {
			Name:        "list",
			Description: "Display current threshold values and profiles.",
			Callback:    commandList,
		},
		"health": {
			Name:        "health",
			Description: "Display percentage of full-charge spec remaining.",
			Callback:    commandHealth,
		},
		"create": {
			Name:        "create",
			Description: "Create or overwrite a charge threshold profile.",
			Callback:    commandCreate,
		},
		"delete": {
			Name:        "delete",
			Description: "Delete an existing charge threshold profile.",
			Callback:    commandDelete,
		},
		"set": {
			Name:        "set",
			Description: "Activate an existing charge threshold profile.",
			Callback:    commandSet,
		},
	}
}
