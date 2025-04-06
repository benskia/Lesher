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

// TODO: Command execution

type Command struct {
	Name string
	Cmd  func(args []string, cfg *config.Config) error
}

func Execute(args []string)
