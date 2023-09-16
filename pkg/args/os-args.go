package args

import (
	"fmt"

	"github.com/mccuskero/sandbox/pkg/config"
)

func NewConfigFromOsArgs(len int, args []string) (*config.Config) {

	fmt.Println("arg len is: ", len)

	if len <= 1 {
		fmt.Println("Need to pass at least on argument")
		return nil
	}

	// using a struct literal
	c := &config.Config{
		AwsRegion: args[1],
	}

	return c
}