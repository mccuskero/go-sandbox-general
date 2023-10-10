package bloom

import (
	"fmt"
	"testing"

	"github.com/pborman/uuid"
)

func TestSimpleBloom(t* testing.T) {
	sb := NewSimpleBloom()

	id1 := uuid.NewRandom().String()
	sb.Set(id1)

	id := uuid.NewRandom().String()
	sb.Set(id)

	found, _ := sb.Get(id1)

	fmt.Println(found)
}