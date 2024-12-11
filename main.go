package main

import (
	"fmt"
	"github.com/dunnwithpuns/gator/internal/config"
)

func main() {

	cfg, err := config.READ()
	if err != nil {
		error := fmt.Errorf("error reading config file: %v", err)
		fmt.Println(error)
	}

	config.SetUser("zack")

	cfg, err = config.READ()
	if err != nil {
		error := fmt.Errorf("read error: %v", err)
		fmt.Println(error)
	}


	fmt.Println(cfg)
}
