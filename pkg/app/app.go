package app

import (
	"fmt"
	"os"
	"time"

	"github.com/fishykins/clam/pkg/cli"
)

type App struct {
	Name        string
	Description string
	TickSpeed   int
	input       *chan cli.Message
	commands    []cli.Command
}

func NewApp(in *chan cli.Message, name string, description string) *App {
	return &App{
		Name:        name,
		Description: description,
		TickSpeed:   1,
		input:       in,
		commands:    []cli.Command{},
	}
}

func (a *App) AddCommand(c cli.Command) {
	a.commands = append(a.commands, c)
}

func (a *App) HandleInput(args []string) {
	if cli.MatchCommand(args, cli.Help) {
		cli.PrintHelp(a.commands)
		return
	}
	command, err := cli.MatchCommands(args, a.commands)
	if err != nil {
		a.Log(err)
		return
	}
	if command.Action != nil {
		err = command.Action(args)
		if err != nil {
			a.Log(err)
		}
	}
}

func (a *App) Run() {
	a.Log("started...")

	if a.input != nil {
		a.Log("Listening for input...")
		go func() {
			for {
				input := <-*a.input
				if input.Sender == "ctrlc" {
					a.Stop()
					return
				}
				if input.Sender == "cli" {
					if len(input.Arguments) > 0 {
						a.HandleInput(input.Arguments)
					}
				}
			}
		}()
	}

	for {
		Tick(a)
		time.Sleep(time.Duration(a.TickSpeed) * time.Millisecond)
	}
}

func (a *App) Stop() {
	a.Log("Stopping app...")
	os.Exit(1)
}

func (a *App) Log(message interface{}) {
	fmt.Printf("%s: %s\n", a.Name, message)
}

func (a *App) Logf(format string, args ...interface{}) {
	format = fmt.Sprintf("%s: %s", a.Name, format)
	fmt.Printf(format, args...)
}
