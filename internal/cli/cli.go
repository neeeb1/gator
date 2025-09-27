package cli

import (
	"context"
	"fmt"

	"github.com/neeeb1/gator/internal/config"
	"github.com/neeeb1/gator/internal/database"
)

type State struct {
	Config *config.Config
	Db     *database.Queries
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

func MiddlewareLoggedIn(handler func(s *State, cmd Command, user database.User) error) func(*State, Command) error {
	return func(s *State, cmd Command) error {

		u, err := s.Db.GetUser(context.Background(), s.Config.CurrentUser)
		if err != nil {
			return err
		}

		return handler(s, cmd, u)
	}
}
