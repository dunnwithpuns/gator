package cli

import (
	"fmt"
)

type command struct {
	name string
	arguments string
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
