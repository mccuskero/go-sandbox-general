package main

import (
	"fmt"
	"os"

	"github.com/mccuskero/go-sandbox-general/pkg/cli"
)

// need to find a way to manage the application/business logic in the cobra Run section.

func main() {
	if err := cli.NewCobraAppCmd("test").Execute(); err != nil {
		//		fmt.Println("Error starting up CobraApp: ", err)
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
	}
}
