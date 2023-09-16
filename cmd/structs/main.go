package main

import (
	"fmt"

	"github.com/mccuskero/sandbox/pkg/structs"
)

func main() {
	fmt.Println("Starting structs sandbox.")

	p, err := structs.NewPerson("John", "NYC")
	if err != nil {
		fmt.Println("Error creating person: ", err)
	}

	structs.PrettyPrint(p)

	t, _ := structs.NewTelephone("mobile", "number")
	c, _ := structs.NewContact("John", t)
	structs.PrettyPrint(c)

	t2, _ := structs.NewTelephone("mobile", "number")
	c.AddOtherNumber(t2)
	structs.PrettyPrint(c)
}
