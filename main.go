package main

import (
	"fmt"
	"github.com/dunnwithpuns/gator/internal/config"
	"github.com/dunnwithpuns/gator/internal/cli"
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

}
