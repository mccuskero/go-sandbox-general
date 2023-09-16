package main

import (
	"fmt"
	"strings"

	"github.com/pborman/uuid"
)

func main() {
	fmt.Println("Hello")

	uuidWithHyphen := uuid.NewRandom()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	fmt.Println(uuid)
}
