package cli

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
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
		pubTime, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			return err
		}

		post_params := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       sql.NullString{String: item.Title, Valid: true},
			Url:         item.Link,
			Description: sql.NullString{String: item.Description, Valid: true},
			PublishedAt: sql.NullTime{
				Time:  pubTime,
				Valid: true,
			},
			FeedID: nextFeed.ID,
		}

		_, err = s.Db.CreatePost(context.Background(), post_params)
		if err != nil && !strings.Contains(err.Error(), `duplicate key value violates unique constraint "posts_url_key"`) {
			fmt.Println(err)
		}
	}
	fmt.Println()

	return nil
}
