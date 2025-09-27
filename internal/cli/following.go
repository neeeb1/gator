package cli

import (
	"context"
	"fmt"

	"github.com/neeeb1/gator/internal/database"
)

func HandlerFollowing(s *State, cmd Command, u database.User) error {
	if len(cmd.Arguments) != 0 {
		return fmt.Errorf("expected no arguments, but got %d", len(cmd.Arguments))
	}

	feeds, err := s.Db.GetFeedFollowsForUser(context.Background(), u.ID)
	if err != nil {
		return err
	}

	fmt.Printf("followed feeds for %s\n", u.Name)
	for _, f := range feeds {
		fmt.Printf("* %s\n", f.FeedName)
	}

	return nil
}
