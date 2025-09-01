package main

import (
	"context"
	"fmt"
	"time"

	"github.com/lukaspodobnik/gator/internal/database"
	"github.com/lukaspodobnik/gator/internal/rss"
)

func aggHandler(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <time_between_reqs>", cmd.name)
	}

	duration, err := time.ParseDuration(cmd.args[0] + "m")
	if err != nil {
		return fmt.Errorf("parsing <time_between_reqs> failed: %w", err)
	}

	fmt.Printf("Collecting feeds every %v\n\n", duration)

	ticker := time.NewTicker(duration)
	for ; ; <-ticker.C {
		fetchOldest(s)
		fmt.Println()
	}
}

func fetchOldest(s *state) {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		fmt.Println("getting next feed to fetch failed")
	}

	content, err := rss.FetchFeed(context.Background(), feed.Url)
	if err != nil {
		fmt.Println("fetching the feed failed")
		return
	}

	if err := s.db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		UpdatedAt: time.Now(),
		ID:        feed.ID,
	}); err != nil {
		fmt.Println("marking feed as fetched failed")
	}

	for _, item := range content.Channel.Item {
		fmt.Println(item.Title)
	}
}
