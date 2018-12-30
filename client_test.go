package push

import (
	"fmt"
	"testing"
)

func TestPull(t *testing.T) {
	userID := 1234
	clID := 123

	s := NewSession(userID)
	c := s.NewClient(clID)
	defer c.DeleteSelf()
	go s.Push("Hello")
	msg, err := c.Pull()
	if err != nil {
		panic(err)
	}

	if m, ok := msg.(string); ok {
		if m == "Hello" {
			fmt.Println(m)
			return
		}
	}
	t.Fail()
}
