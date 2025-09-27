package cli

import (
	"context"
	"fmt"
	"strconv"

	"github.com/neeeb1/gator/internal/database"
)

func HandlerBrowse(s *State, cmd Command, u database.User) error {
	limit := 2
	var err error

	if len(cmd.Arguments) > 1 {
		return fmt.Errorf("expected 1 arguments, but got %d", len(cmd.Arguments))
	}
	if len(cmd.Arguments) == 1 {
		limit, err = strconv.Atoi(cmd.Arguments[0])
		if err != nil {
			return err
		}
	}

	params := database.GetPostsForUserParams{
		Limit:  int32(limit),
		UserID: u.ID,
	}

	posts, err := s.Db.GetPostsForUser(context.Background(), params)
	if err != nil {
		return err
	}

	fmt.Printf("Getting the newest %d posts for %s", limit, u.Name)
	for i, p := range posts {
		feed, err := s.Db.GetFeedById(context.Background(), p.FeedID)
		if err != nil {
			return err
		}

		fmt.Printf(`
			%d.) %s
			from %s
			Published: %s
			Link: %s

		`, i+1, p.Title.String, feed.Name, p.PublishedAt.Time, p.Url)
	}

	return nil
}
