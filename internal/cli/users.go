package cli

import (
	"context"
	"fmt"
)

func HandlerUsers(s *State, cmd Command) error {
	if len(cmd.Arguments) > 0 {
		return fmt.Errorf("expected no arguments, but got %d", len(cmd.Arguments))
	}

	users, err := s.Db.GetAllUsers(context.Background())
	if err != nil {
		return err
	}

	for _, u := range users {
		if u.Name == s.Config.CurrentUser {
			fmt.Printf("* %s (current)\n", u.Name)
		} else {
			fmt.Printf("* %s\n", u.Name)
		}
	}

	return nil
}
