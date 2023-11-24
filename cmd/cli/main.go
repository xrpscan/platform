package main

import (
	"errors"
	"fmt"
	"os"
)

// Command runner interface
type Runner interface {
	Init([]string) error
	Validate() error
	Name() string
	Run() error
}

// command line root
func root(args []string) error {
	if len(args) < 1 {
		return errors.New("you must pass a sub command")
	}
	subcommand := os.Args[1]

	cmds := []Runner{
		NewBackfillCommand(),
		NewInitCommand(),
	}

	for _, cmd := range cmds {
		if cmd.Name() == subcommand {
			cmd.Init(os.Args[2:])
			if err := cmd.Validate(); err != nil {
				return err
			}

			return cmd.Run()
		}
	}

	return fmt.Errorf("unknown subcommand: %s", subcommand)
}

func main() {
	if err := root(os.Args[1:]); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
