package push

import (
	"fmt"
	"testing"
)

func TestPush(t *testing.T) {
	userID := 1234
	clID := 123

	s := NewSession(userID)
	c := s.NewClient(clID)
	defer c.DeleteSelf()

	ch, err := c.PullChan()
	if err != nil {
		panic(err)
	}

	go s.Push("Hello")
	msg := <-ch

	if m, ok := msg.(string); ok {
		if m == "Hello" {
			fmt.Println(msg)
			return
		}
	}
	t.Fail()
}
