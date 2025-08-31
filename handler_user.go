package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/lukaspodobnik/gator/internal/database"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.name)
	}

	name := cmd.args[0]
	if _, err := s.db.GetUser(context.Background(), name); err != nil {
		return fmt.Errorf("user %s does not exist: %w", name, err)
	}

	err := s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User switched successfully!")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.name)
	}

	name := cmd.args[0]
	if _, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
	}); err != nil {
		return fmt.Errorf("creating user in db failed: %w", err)

	}

	if err := s.cfg.SetUser(name); err != nil {
		return fmt.Errorf("setting user in cfg failed: %w", err)
	}

	fmt.Printf("User %s successfully registered!", name)
	return nil
}

func handlerReset(s *state, cmd command) error {
	if err := s.db.DeleteAllUsers(context.Background()); err != nil {
		return err
	}

	fmt.Println("All rows in users successfully deleted!")
	return nil
}
