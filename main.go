package main

import (
	"fmt"

	"github.com/neeeb1/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Errorf("error reading config: %v", err)
	}

	cfg.SetUser("neeeb")

	cfg, err = config.Read()
	if err != nil {
		fmt.Errorf("error reading config: %v", err)
	}
	fmt.Println(cfg)
}
