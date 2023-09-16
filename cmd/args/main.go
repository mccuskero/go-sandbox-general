package main

import (
	"os"

	"github.com/mccuskero/sandbox/pkg/args"
	"github.com/mccuskero/sandbox/pkg/structs"
)

func main() {
	c := args.NewConfigFromOsArgs(len(os.Args), os.Args)
	structs.PrettyPrint(c)
}