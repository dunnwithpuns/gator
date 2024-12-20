package cli

import (
	"fmt"
)

type Commands struct {
	CommandsMap map[string]func(*State, Command) error
}

func (c *Commands) Register(name string, f func(*State, Command) error) {

	c.CommandsMap[name] = f

}

func (c *Commands) Run(s *State, cmd Command) error {

	handlerFunction, exists := c.CommandsMap[cmd.Name]
	if !exists {
		return fmt.Errorf("%v is not a registered command", cmd)
	}

	return handlerFunction(s, cmd)
}
