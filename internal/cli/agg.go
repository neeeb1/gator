package cli

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/neeeb1/gator/internal/database"
	"github.com/neeeb1/gator/internal/rss"
)

func HandlerAgg(s *State, cmd Command) error {
	if len(cmd.Arguments) != 1 {
		return fmt.Errorf("expected 1 arguments, but got %d", len(cmd.Arguments))
	}

	freq, err := time.ParseDuration(cmd.Arguments[0])
	if err != nil {
		return err
	}

	ticker := time.NewTicker(freq)
	fmt.Printf("fetching feeds every %s\n\n", freq)

	for ; ; <-ticker.C {
		err = scrapeFeeds(s)
		if err != nil {
			return err
		}
	}
}

func scrapeFeeds(s *State) error {
	nextFeed, err := s.Db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	params := database.MarkFeedFetchedParams{
		LastFetchedAt: sql.NullTime{Time: time.Now(), Valid: true},
		ID:            nextFeed.ID,
	}

	_, err = s.Db.MarkFeedFetched(context.Background(), params)
	if err != nil {
		return err
	}

	fetched, err := rss.FetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return err
	}

	fmt.Printf("Fetching rss feed %s at %s\n", nextFeed.Name, nextFeed.Url)
	for _, item := range fetched.Channel.Item {
		fmt.Printf("* %s\n", item.Title)
	}
	fmt.Println()

	return nil
}
