package push

import "testing"

func TestPush(t *testing.T) {
	c := NewClient("userID", "UniqueClientID")
	c.Push("message")
	m, err := c.Pull("UniqueClientID")
	if err != nil {
		panic(err)
	}
	if message, ok := m.(string); ok {
		if message == "message" {
			return
		}
	}
	t.Fail()
}
