package structs

import (
	"errors"
)

type Telephone struct {
	Number string
	Type   string
}

type Contact struct {
	Name      string
	MainTel   Telephone
	OtherTels []Telephone
}

func NewTelephone(telType string, number string) (*Telephone, error) {
	if telType == "" || number == "" {
		return nil, errors.New("need both telephone type, and number to create telephone")
	}

	t := &Telephone{
		Number: number,
		Type:   telType,
	}

	return t, nil
}

func NewContact(Name string, tel *Telephone) (*Contact, error) {

	if Name == "" {
		return nil, errors.New("name cannot be empty")
	}

	if tel == nil {
		return nil, errors.New("telephone cannot be empty")
	}

	c := new(Contact)

	c.Name = Name
	c.MainTel.Type = tel.Type
	c.MainTel.Number = tel.Number

	return c, nil
}

func (c *Contact) AddOtherNumber(tel *Telephone) error {
	if tel == nil {
		return errors.New("other telephone cannot be nil")
	}

	c.OtherTels = append(c.OtherTels, *tel)

	return nil
}
