package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/lukaspodobnik/gator/internal/database"
)

func followHandler(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.name)
	}

	url := cmd.args[0]

	feed, err := s.db.GetFeed(context.Background(), url)
	if err != nil {
		return fmt.Errorf("getting feed %s failed: %w", url, err)
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})

	if err != nil {
		return fmt.Errorf("creating feed-follow failed: %w", err)
	}

	fmt.Printf("%s now follows %s\n", feedFollow.UserName, feedFollow.FeedName)
	return nil
}

func followingHandler(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("usage: %s", cmd.name)
	}

	follows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("getting feed-follows for %s failed: %w", user.Name, err)
	}

	fmt.Printf("%s is following:\n", user.Name)
	for _, follow := range follows {
		fmt.Printf("  - %s\n", follow.FeedName)
	}

	return nil
}

func unfollowHandler(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.name)
	}

	url := cmd.args[0]
	feed, err := s.db.GetFeed(context.Background(), url)
	if err != nil {
		return fmt.Errorf("getting feed failed: %w", err)
	}

	if err := s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}); err != nil {
		return fmt.Errorf("deleting feed-follow failed: %w", err)
	}

	fmt.Println("Successfully unfollowed!")
	return nil
}
