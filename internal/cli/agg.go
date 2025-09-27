package cli

import (
	"context"
	"fmt"

	"github.com/neeeb1/gator/internal/rss"
)

func HandlerAgg(s *State, cmd Command) error {
	if len(cmd.Arguments) > 0 {
		return fmt.Errorf("expected no arguments, but got %d", len(cmd.Arguments))
	}

	fmt.Println(rss.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml"))
	return nil
}
