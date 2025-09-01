package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/lukaspodobnik/gator/internal/database"
)

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

	if _, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    current_user.ID,
		FeedID:    feed.ID,
	}); err != nil {
		return fmt.Errorf("could not follow the created feed: %w", err)
	}

	fmt.Printf("%s is now following %s\n", current_user.Name, feed.Name)
	return nil
}

func feedsHandler(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("usage: %s", cmd.name)
	}

	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("getting feeds failed: %w", err)
	}

	for _, feed := range feeds {
		fmt.Printf("feed name: %s, feed url: %s, created by: %s\n", feed.Name, feed.Url, feed.CreatedBy)
	}

	return nil
}
