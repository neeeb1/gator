package cli

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/neeeb1/gator/internal/database"
)

func HandlerFollow(s *State, cmd Command, u database.User) error {
	if len(cmd.Arguments) != 1 {
		return fmt.Errorf("expected 1 arguments, but got %d", len(cmd.Arguments))
	}

	f, err := s.Db.GetFeed(context.Background(), cmd.Arguments[0])
	if err != nil {
		return err
	}

	params := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    u.ID,
		FeedID:    f.ID,
	}

	_, err = s.Db.CreateFeedFollow(context.Background(), params)
	if err != nil {
		return err
	}

	fmt.Println("feed successfully followed")
	fmt.Printf("feed: %s, user: %s", f.Name, u.Name)

	return nil
}
