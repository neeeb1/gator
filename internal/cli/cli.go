package cli

import (
	"fmt"

	"github.com/neeeb1/gator/internal/config"
)

type State struct {
	Config config.Config
}

type Command struct {
	Name      string
	Arguments []string
}

type Commands struct {
	Cmds map[string]func(*State, Command) error
}

func (c *Commands) Run(s *State, cmd Command) error {
	//fmt.Printf("Attempting to run command %s", cmd.Name)

	if _, exists := c.Cmds[cmd.Name]; !exists {
		return fmt.Errorf("Command does not exist")
	}

	err := c.Cmds[cmd.Name](s, cmd)
	if err != nil {
		return err
	}
	return nil
}

func (c *Commands) Register(name string, f func(*State, Command) error) error {
	c.Cmds[name] = f
	return nil
}
