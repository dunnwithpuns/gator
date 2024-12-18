package cli

import (
	"fmt"
)

type Command struct {
	Name string
	Arguments []string
}

func LoginHandler(s *State, cmd Command) error {

	if len(cmd.Arguments) == 0 {
		return fmt.Errorf("No username submitted")
	} else if len(cmd.Arguments) > 1 {
		return fmt.Errorf("Too many arguments submitted")
	}

	err := s.Config.SetUser(cmd.Arguments[0])
	if err != nil {
		return err
	}

	fmt.Printf("The user name has been set to: %v", cmd.Arguments)

	return nil
}
