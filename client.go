package push

// Client holds the methods for a perticular client.
type Client struct {
	clientID interface{}
	session  *Session
}

// Pull pulls a message for the current client.
/*
	var (
		userID := 123446555
		clID := "device_mobile_5445"
	)

	s := NewSession(userID)
	c := s.NewClient(clID)
	defer c.DeleteSelf()

	msg, err := c.Pull()
	if err != nil {
		panic(err)
	}
	fmt.Println(msg)
*/
func (c *Client) Pull() (content interface{}, err error) {
	return c.session.pull(c.clientID)
}

// PullChan returns a channel for receiving messages for the current client.
/*
	var (
		userID := 123446555
		clID := "device_mobile_5645"
	)

	s := NewSession(userID)
	c := s.NewClient(clID)
	defer c.DeleteSelf()

	ch, err := c.PullChan()
	if err != nil {
		panic(err)
	}

	msg := <-ch
	fmt.Println(msg)

*/
//
// Exclusively usable with websockets
func (c *Client) PullChan() (ClientChan, error) {
	return c.session.pullChan(c.clientID)
}

// Key returns the current clientID/name/key.
func (c *Client) Key() interface{} {
	return c.clientID
}

// DeleteSelf deletes the current client from the current session.
func (c *Client) DeleteSelf() {
	delete(c.session.clients, c.clientID)
}
