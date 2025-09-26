package cli

import (
	"fmt"
)

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Arguments) == 0 {
		return fmt.Errorf("no arguments found")
	} else if len(cmd.Arguments) > 1 {
		return fmt.Errorf("expected 1 argument, but recieved %d", len(cmd.Arguments))
	}

	err := s.Config.SetUser(cmd.Arguments[0])
	if err != nil {
		return err
	}

	fmt.Printf("%s has been set as active user\n", cmd.Arguments[0])
	return nil
}
