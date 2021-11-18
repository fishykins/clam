package cli

import "fmt"

var Help = Command{
	Name:        "help",
	Description: "Prints this help message",
	Aliases:     []string{"h"},
}

func PrintHelp(commands []Command) {
	fmt.Println("Available commands:")
	for _, command := range commands {
		fmt.Printf("%s - %s\n", command.Name, command.Description)
	}
}
