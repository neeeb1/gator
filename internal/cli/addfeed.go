package cli

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/neeeb1/gator/internal/database"
)

func HandlerAddFeed(s *State, cmd Command) error {
	if len(cmd.Arguments) == 0 {
		return fmt.Errorf("no arguments found")
	}
	if len(cmd.Arguments) != 2 {
		return fmt.Errorf("expected 2 argument, but recieved %d", len(cmd.Arguments))
	}

	u, err := s.Db.GetUser(context.Background(), s.Config.CurrentUser)
	if err != nil {
		return err
	}

	params := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Arguments[0],
		Url:       cmd.Arguments[1],
		UserID:    u.ID,
	}

	feed, err := s.Db.CreateFeed(context.Background(), params)
	if err != nil {
		return err
	}

	follow_params := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    u.ID,
		FeedID:    params.ID,
	}

	_, err = s.Db.CreateFeedFollow(context.Background(), follow_params)
	if err != nil {
		return err
	}

	fmt.Println("new feed added")
	fmt.Printf("feed name: %s\ncreated at: %s\nupdated at: %s\nfeed uuid: %s\n",
		feed.Name,
		feed.CreatedAt,
		feed.UpdatedAt,
		feed.ID,
	)

	return nil
}
