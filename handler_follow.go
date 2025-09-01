package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/lukaspodobnik/gator/internal/database"
)

func followHandler(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.name)
	}

	url := cmd.args[0]

	current_user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("getting user %s failed: %w", s.cfg.CurrentUserName, err)
	}

	feed, err := s.db.GetFeed(context.Background(), url)
	if err != nil {
		return fmt.Errorf("getting feed %s failed: %w", url, err)
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    current_user.ID,
		FeedID:    feed.ID,
	})

	if err != nil {
		return fmt.Errorf("creating feed-follow failed: %w", err)
	}

	fmt.Printf("%s now follows %s\n", feedFollow.UserName, feedFollow.FeedName)
	return nil
}

func followingHandler(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("usage: %s", cmd.name)
	}

	currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("getting user %s failed: %w", s.cfg.CurrentUserName, err)
	}

	follows, err := s.db.GetFeedFollowsForUser(context.Background(), currentUser.ID)
	if err != nil {
		return fmt.Errorf("getting feed-follows for %s failed: %w", currentUser.Name, err)
	}

	fmt.Printf("%s is following:\n", currentUser.Name)
	for _, follow := range follows {
		fmt.Printf("  - %s\n", follow.FeedName)
	}

	return nil
}
