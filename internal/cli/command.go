package cli

import (
	"fmt"
)

type Command struct {
	name string
	arguments string
}

func LoginHandler(s *State, cmd Command) error {

	if cmd.arguments == "" {
		return fmt.Errorf("No username submitted")
	}

	err := s.Config.SetUser(cmd.arguments)
	if err != nil {
		return err
	}

	fmt.Printf("The user name has been set to: %v", cmd.arguments)

	return nil
}