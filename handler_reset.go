package main

import (
	"context"
	"fmt"
)

func resetHandler(s *state, cmd command) error {
	if err := s.db.DeleteAllUsers(context.Background()); err != nil {
		return err
	}

	fmt.Println("All rows in users successfully deleted!")
	return nil
}
