package cli

import (
	"fmt"
	"os"
)

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Arguments) == 0 {
		defer os.Exit(1)
		return fmt.Errorf("no arguments found")
	} else if len(cmd.Arguments) > 1 {
		defer os.Exit(1)
		return fmt.Errorf("expected 1 argument, but recieved %d", len(cmd.Arguments))
	}

	err := s.Config.SetUser(cmd.Arguments[0])
	if err != nil {
		return err
	}

	fmt.Printf("%s has been set as active user", cmd.Arguments[0])
	return nil
}
