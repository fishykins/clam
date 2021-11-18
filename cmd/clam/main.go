package main

import (
	"fmt"

	"github.com/fishykins/clam/pkg/app"
	"github.com/fishykins/clam/pkg/cli"
)

func bunnies(args []string) error {
	fmt.Println("I love bunnies, bunnies are fluffy and cute.")
	return nil
}

func main() {

	// Create app and input channels
	io := make(chan cli.Message)
	app := app.NewApp(&io, "Clam", "Test application")

	// Add commands
	app.AddCommand(cli.Command{Name: "bunny", Description: "bunnies are nice", Aliases: []string{"bunnies"}, Action: bunnies})

	// Start input loop
	go cli.HandleCtrlC(&io)
	go cli.ConsoleInput(&io)

	app.Run()
}
