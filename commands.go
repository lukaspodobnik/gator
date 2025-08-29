package main

import (
	"errors"
)

type command struct {
	name string
	args []string
}

type commands struct {
	handlers map[string]func(*state, command) error
}

func (c *commands) register(name string, handler func(*state, command) error) {
	c.handlers[name] = handler
}

func (c *commands) run(s *state, cmd command) error {
	handler, exits := c.handlers[cmd.name]
	if !exits {
		return errors.New("no handler registered for given command ")
	}

	return handler(s, cmd)
}
