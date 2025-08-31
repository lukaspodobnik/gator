package main

import (
	"context"
	"fmt"

	"github.com/lukaspodobnik/gator/internal/rss"
)

func aggHandler(s *state, cmd command) error {
	url := "https://www.wagslane.dev/index.xml"
	feed, err := rss.FetchFeed(context.Background(), url)
	if err != nil {
		return fmt.Errorf("fetiching feed at %s failed: %w", url, err)
	}

	fmt.Println(feed)

	return nil
}
