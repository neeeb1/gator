package cli

import (
	"context"
	"fmt"
)

func HandlerReset(s *State, cmd Command) error {
	if len(cmd.Arguments) > 0 {
		return fmt.Errorf("expected no arguments, but got %d", len(cmd.Arguments))
	}

	err := s.Db.ResetUsers(context.Background())
	if err != nil {
		return err
	}

	return nil
}
