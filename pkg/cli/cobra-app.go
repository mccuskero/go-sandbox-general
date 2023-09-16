package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version = "0.0.1"

func NewCobraAppCmd(in string) *cobra.Command {
	c := &cobra.Command{
		Use:   "cobraApp",
		Version: version,
		Short: "cobraApp is a very fast CLI tool",
		Long: `A Fast and Flexible CLI Tools with
				  love by spf13 and friends in Go.
				  Complete documentation is available at http://hugo.spf13.com`,
		RunE: func(cmd *cobra.Command, args []string) (error) {
			// probably want to import the app logic here, and execute it... pass in the cobra comand and args... 
//			fmt.Fprintf(cmd.OutOrStdout(), in)
			term, _ := cmd.Flags().GetString("term")
			fmt.Println("The term passed in was:", term)

			return nil
		},
	}

	c.PersistentFlags().String("term", "", "A search term for a dad joke.")

	return c
}

