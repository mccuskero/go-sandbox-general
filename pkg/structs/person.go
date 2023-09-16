package structs

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Person struct {
	Name string
	City string
}

func NewPerson(name string, city string) (*Person, error) {

	if name == "" || city == "" {
		return nil, errors.New("cannot create person with empty name, or city string")
	}

	p := new(Person)
	p.Name = name
	p.City = city

	return p, nil
}

func PrettyPrint(v interface{}) {
	s, _ := json.MarshalIndent(v, "", "\t")
	fmt.Println(string(s))
}
