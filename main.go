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
	}

	state := cli.State{
		Config: &cfg,
	}

	commands := cli.Commands{
		CommandsMap: make(map[string]func(*cli.State, cli.Command) error),
	}

	commands.Register("login", cli.LoginHandler)

	var input string

	if len(os.Args) > 2 {
		fmt.Printf("Not enough arguments entered")
	} else if len(os.Args) == 2{
		
	}


}
