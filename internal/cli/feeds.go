package cli

import (
	"context"
	"fmt"
)

func HandlerFeeds(s *State, cmd Command) error {
	if len(cmd.Arguments) > 0 {
		return fmt.Errorf("expected no arguments, but got %d", len(cmd.Arguments))
	}

	feeds, err := s.Db.GetAllFeeds(context.Background())
	if err != nil {
		return err
	}

	for _, f := range feeds {
		fmt.Printf("* name: %s, url: %s, created by: %s \n", f.Name, f.Url, f.CreatedBy)
	}

	return nil
}
