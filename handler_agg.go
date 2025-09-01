package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
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
		if err := s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			Title:       item.Title,
			Url:         item.Link,
			Description: item.Description,
			PublishedAt: parseTime(item.PubDate),
			FeedID:      feed.ID,
		}); err != nil {
			fmt.Printf("creating post failed: %v\n", err)
		}
	}
}

func parseTime(t string) time.Time {
	var err error
	publishedAt := time.Now()
	for _, layout := range []string{
		"2006-01-02T15:04:05Z07:00",
		"2006-01-02T15:04:05.000Z07:00",
		"Mon, 02 Jan 2006 15:04:05 -0700",
		"Mon, 02 Jan 2006 15:04:05 MST",
		"02 Jan 06 15:04 -0700",
		"2006-01-02 15:04:05",
		"2006-01-02",
	} {
		publishedAt, err = time.Parse(layout, t)
		if err == nil {
			fmt.Println("parsing pubdate worked")
			break
		}
	}
	return publishedAt.UTC()
}
