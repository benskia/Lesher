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
			Description: helpDescription,
			Callback:    commandHelp,
		},
		"list": {
			Name:        "list",
			Description: listDescription,
			Callback:    commandList,
		},
		"health": {
			Name:        "health",
			Description: healthDescription,
			Callback:    commandHealth,
		},
		"create": {
			Name:        "create",
			Description: createDescription,
			Callback:    commandCreate,
		},
		"delete": {
			Name:        "delete",
			Description: deleteDescription,
			Callback:    commandDelete,
		},
		"set": {
			Name:        "set",
			Description: setDescription,
			Callback:    commandSet,
		},
	}
}
