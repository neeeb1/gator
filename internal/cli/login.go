package cli

import (
	"context"
	"fmt"
)

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Arguments) == 0 {
		return fmt.Errorf("no arguments found")
	} else if len(cmd.Arguments) > 1 {
		return fmt.Errorf("expected 1 argument, but recieved %d", len(cmd.Arguments))
	}

	user, err := s.Db.GetUser(context.Background(), cmd.Arguments[0])
	if err != nil {
		return err
	}

	err = s.Config.SetUser(user.Name)
	if err != nil {
		return err
	}

	fmt.Printf("%s has been set as active user\n", cmd.Arguments[0])
	return nil
}
