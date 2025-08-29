package main

import (
	"log"
	"os"

	"github.com/lukaspodobnik/gator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	s := &state{
		cfg: &cfg,
	}

	commands := commands{
		handlers: make(map[string]func(*state, command) error),
	}
	commands.register("login", handlerLogin)

	args := os.Args
	if len(args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	if err := commands.run(s, command{name: args[1], args: args[2:]}); err != nil {
		log.Fatal(err)
	}
}
