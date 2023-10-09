package measure

import (
	"testing"
	"time"
)

func TestDuration(t* testing.T) {
	defer Duration(Track("TestDuration"))
	time.Sleep(2000000000)  // 2 seconds 
}