package cli

import (
	"fmt"
)

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
