package cli

import (
	"context"
	"fmt"

	"github.com/neeeb1/gator/internal/database"
)

func HandlerUnfollow(s *State, cmd Command, u database.User) error {
	if len(cmd.Arguments) != 1 {
		return fmt.Errorf("expected 1 arguments, but got %d", len(cmd.Arguments))
	}

	f, err := s.Db.GetFeed(context.Background(), cmd.Arguments[0])
	if err != nil {
		return err
	}

	params := database.DeleteFeedFollowParams{
		UserID: u.ID,
		FeedID: f.ID,
	}

	err = s.Db.DeleteFeedFollow(context.Background(), params)
	if err != nil {
		return err
	}

	fmt.Printf("successfully unfollowed %s for user %s", f.Name, u.Name)

	return nil
}
