package main

import (
	"os"

	"github.com/mccuskero/go-sandbox-general/pkg/args"
	"github.com/mccuskero/go-sandbox-general/pkg/structs"
)

func main() {
	c := args.NewConfigFromOsArgs(len(os.Args), os.Args)
	structs.PrettyPrint(c)
}
