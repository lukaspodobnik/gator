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
	url := "https://www.wagslane.dev/index.xml"
	feed, err := rss.FetchFeed(context.Background(), url)
	if err != nil {
		return fmt.Errorf("fetiching feed at %s failed: %w", url, err)
	}

	fmt.Println(feed)

	return nil
}

func addfeedHandler(s *state, cmd command) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.name)
	}

	name := cmd.args[0]
	url := cmd.args[1]

	current_user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("getting user %s failed: %w", s.cfg.CurrentUserName, err)
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    current_user.ID,
	})
	if err != nil {
		return fmt.Errorf("creating feed in db failed: %w", err)
	}

	fmt.Println("Adding feed was successful!")
	fmt.Println(feed)
	return nil
}
