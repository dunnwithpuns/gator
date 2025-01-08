package cli

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dunnwithpuns/gator/internal/database"
	"github.com/google/uuid"
)

type Command struct {
	Name string
	Arguments []string
}

func LoginHandler(s *State, cmd Command) error {

	// verifying number of inputs
	if len(cmd.Arguments) == 0 {
		return fmt.Errorf("No username submitted\n")
	} else if len(cmd.Arguments) > 1 {
		return fmt.Errorf("Too many arguments submitted\n")
	}

	name := cmd.Arguments[0]

	ctx := context.Background()

	_, err := s.db.GetUser(ctx, name)
	if errors.Is(err, sql.ErrNoRows) {
		fmt.Printf("Requested username does not exist\n")
		return err
	}
	if err != nil {
		return err
	}

	// setting the current user to the input user
	err = s.cfg.SetUser(name)
	if err != nil {
		return err
	}

	fmt.Printf("The user name has been set to: %v \n", name)

	return nil
}

func RegisterHandler(s *State, cmd Command) error {

	// verifying that only one input was submitted
	if len(cmd.Arguments) == 0 {
		return fmt.Errorf("No user submitted\n")
	} else if len(cmd.Arguments) > 1 {
		return fmt.Errorf("Too many arugments submitted\n")
	}

	name := cmd.Arguments[0]

	params := database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: name,
	}

	// checking if user exists
	_, err := s.db.GetUser(context.Background(), name)
	// if err == nil, that user already exists
	if err == nil {
		fmt.Printf("Could not register user because %v already exists", name)
		os.Exit(1)
	}
	// if err is not a sql.ErrNoRows error, return err as it is a different failure condition
	if !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	user, err := s.db.CreateUser(context.Background(), params)
	if err != nil {
		return err
	}

	err = s.cfg.SetUser(name)
	if err != nil {
		return err
	}

	log.Printf("User created: %+v\n", user)

	return nil
}

func ResetHandler(s *State, cmd Command) error {

	if len(cmd.Arguments) > 0 {
		return fmt.Errorf("Reset does not require any arguments")
	}

	err := s.db.ResetUsers(context.Background())
	if err != nil {
		fmt.Printf("Error resetting Users table: %v", err)
		os.Exit(1)
	}

	log.Printf("Users table has been reset")

	return nil
}

func UsersHandler(s *State, cmd Command) error {

	if len(cmd.Arguments) > 0 {
		return fmt.Errorf("Users does not require any arguments")
	}


	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		fmt.Printf("Error retrieving user data: %v", err)
	}

	if len(users) == 0 {
		fmt.Println("No users registered")
		return nil
	}

	currentUser := s.cfg.Current_User_Name

	for _, user := range(users) {
		if user.Name == currentUser {
			fmt.Printf("* %v (current)\n", user.Name)
		} else {
			 fmt.Printf("* %v\n", user.Name)
		}
	}


	return nil
}
