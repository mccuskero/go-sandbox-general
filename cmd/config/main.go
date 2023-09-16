package main

import (
	"fmt"

	"github.com/mccuskero/sandbox/pkg/config"
)

func main() {
	fmt.Println("Starting config sandbox.")

	cfg, err := config.NewFromRegion("test")

	if err != nil {
		fmt.Println("Error creating config: ", err)
	}

	fmt.Println("Config: ", cfg)
}
