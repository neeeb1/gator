package main

import (
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/neeeb1/gator/internal/cli"
	"github.com/neeeb1/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("error reading config: %v\n", err)
		return
	}

	state := cli.State{
		Config: &cfg,
	}

	commands := cli.Commands{
		Cmds: make(map[string]func(*cli.State, cli.Command) error),
	}
	commands.Register("login", cli.HandlerLogin)

	args := os.Args
	if len(args) < 2 {
		fmt.Printf("error: not enough arguments\n")
		os.Exit(1)
	}

	cmd := cli.Command{
		Name:      args[1],
		Arguments: args[2:],
	}

	err = commands.Run(&state, cmd)
	if err != nil {
		fmt.Printf("error running specified command: %v\n", err)
		os.Exit(1)
	}
}
