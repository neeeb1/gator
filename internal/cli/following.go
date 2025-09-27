package cli

import (
	"context"
	"fmt"
)

func HandlerFollowing(s *State, cmd Command) error {
	if len(cmd.Arguments) != 0 {
		return fmt.Errorf("expected no arguments, but got %d", len(cmd.Arguments))
	}

	u, err := s.Db.GetUser(context.Background(), s.Config.CurrentUser)
	if err != nil {
		return err
	}

	feeds, err := s.Db.GetFeedFollowsForUser(context.Background(), u.ID)
	if err != nil {
		return err
	}

	for _, f := range feeds {
		fmt.Printf("followed feeds for %s\n", u.Name)
		fmt.Printf("* %s\n", f.FeedName)
	}

	return nil
}
