package main

import (
	"fmt"
	"github.com/dunnwithpuns/gator/internal/config"
	"github.com/dunnwithpuns/gator/internal/cli"
	"os"
)

func main() {

	cfg, err := config.READ()
	if err != nil {
		error := fmt.Errorf("error reading config file: %v", err)
		fmt.Println(error)
		os.Exit(1)
	}

	state := cli.State{
		Config: &cfg,
	}

	commands := cli.Commands{
		CommandsMap: make(map[string]func(*cli.State, cli.Command) error),
	}

	commands.Register("login", cli.LoginHandler)

	input := os.Args

	if len(input) > 2 {
		fmt.Println("Not enough arguments entered")
		os.Exit(1)
	}

	var inputArguments []string
	
	if len(input) > 3 {
		inputArguments = input[2:]
	}

	inputCommand := cli.Command{
		Name: input[1],
		Arguments: inputArguments,
	}

	err = commands.Run(&state, inputCommand)
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

}
