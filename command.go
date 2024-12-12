package main

import (
	"fmt"
)

type command struct {
	name string
	arguments string
}

type commands struct {
	commandsMap map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {

	c.commandsMap[name] = f

}

func (c *commands) run(s *state, cmd command) error {

	handlerFunction, exists := c.commandsMap[cmd.name]
	if !exists {
		return fmt.Errorf("%v is not a registered command", cmd)
	}

	return handlerFunction(s, cmd)
}

func handlerLogin(s *state, cmd command) error {

	if cmd.arguments == "" {
		return fmt.Errorf("No username submitted")
	}

	err := s.config.SetUser(cmd.arguments)
	if err != nil {
		return err
	}

	fmt.Printf("The user name has been set to: %v", cmd.arguments)

	return nil
}
