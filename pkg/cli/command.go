package cli

import "errors"

type Command struct {
	Name        string
	Description string
	Aliases     []string
	Action      func([]string) error
}

func MatchCommand(args []string, command Command) bool {
	if command.Name == args[0] {
		return true
	}

	// Command not found so search alliases for a match
	for _, alias := range command.Aliases {
		if alias == args[0] {
			return true
		}
	}
	return false
}

func MatchCommands(args []string, commands []Command) (*Command, error) {
	for _, command := range commands {
		if command.Name == args[0] {
			return &command, nil
		}
	}

	// Command not found so search alliases for a match
	for _, command := range commands {
		for _, alias := range command.Aliases {
			if alias == args[0] {
				return &command, nil
			}
		}
	}
	return nil, errors.New("no command found")
}
