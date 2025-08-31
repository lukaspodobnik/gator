package main

// connection string: "postgres://lukaspodobnik:@localhost:5432/gator?sslmode=disable"

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/lukaspodobnik/gator/internal/config"
	"github.com/lukaspodobnik/gator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("error opening db: %v", err)
	}

	dbQueries := database.New(db)
	s := &state{
		db:  dbQueries,
		cfg: &cfg,
	}

	commands := commands{
		handlers: make(map[string]func(*state, command) error),
	}
	commands.register("login", handlerLogin)
	commands.register("register", handlerRegister)
	commands.register("reset", handlerReset)
	commands.register("users", usersHandler)
	commands.register("agg", aggHandler)
	commands.register("addfeed", addfeedHandler)
	commands.register("feeds", feedsHandler)

	args := os.Args
	if len(args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	command := command{name: args[1], args: args[2:]}
	if err := commands.run(s, command); err != nil {
		log.Fatalf("command %s failed: %v", command.name, err)
	}
}
