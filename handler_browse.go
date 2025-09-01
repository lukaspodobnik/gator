package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/lukaspodobnik/gator/internal/database"
)

func browseHandler(s *state, cmd command, user database.User) error {
	var err error
	limit := 2
	if len(cmd.args) > 1 {
		return fmt.Errorf("usage: %s [OPTIONAL] <limit>", cmd.name)
	} else if len(cmd.args) == 1 {
		limit, err = strconv.Atoi(cmd.args[0])
		if err != nil {
			return fmt.Errorf("parsing limit failed: %w", err)
		}
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return fmt.Errorf("getting posts failed: %w", err)
	}

	for _, post := range posts {
		fmt.Printf("Title:       %s\n", post.Title)
		fmt.Printf("Description: %s\n", post.Description)
		fmt.Println()
	}

	return nil
}
