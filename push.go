// package push gives you the ability to use push and pull mechanism for notification or message via websocket or even http client.
package push

import (
	"errors"
)

// NewClient returns a client interface.
//	`sessionID` means the userID or a groupID.
//	`clientID` means the deviceID.
//
// A single user (sessionID) can use multiple devices (clientID).
// That's why the clientID should be unique for each device.
func NewClient(sessionID, clientID interface{}) *Client {
	ch := make(chan interface{})
	if _, ok := cmap[sessionID]; !ok {
		cmap[sessionID] = &Client{
			clients: make(map[interface{}]chan interface{}),
		}
	}
	if clientID != nil {
		cmap[sessionID].clients[clientID] = ch
	}
	return cmap[sessionID]
}

type (
	// Client holds the interface for push and pull mechanism
	Client struct {
		clients map[interface{}]chan interface{}
	}
	clientmap map[interface{}]*Client
)

var cmap = make(clientmap)

// Push sends a message to all connected clients.
/*
	c := push.NewClient(userID, nil)
	c.Push("Hello world!")
*/
func (c *Client) Push(message interface{}) {
	for _, v := range c.clients {
		go func(ch chan interface{}) {
			ch <- message
		}(v)
	}
}

// Pull pulls a message for the given clientID.
/*
	c := push.NewClient(userID, clID)
	for {
		msg, err := c.Pull(clID):
		if err != nil {
			panic(err)
		}
		fmt.Println(msg)
	}
*/
func (c *Client) Pull(clientID interface{}) (content interface{}, err error) {
	if ch, ok := c.clients[clientID]; ok {
		content = <-ch
	} else {
		err = errors.New("push: no such client")
	}
	return
}

// PullChan returns a channel for receiving messages for the given clientID
/*
	c := push.NewClient(userID, clID)
	ch, err := c.PullChan(clID):
	if err != nil {
		panic(err)
	}
	for {
	  select {
		case msg := <- ch:
		fmt.Println(msg)
	  }
	}

*/
//
// Extremely usable with websockets
func (c *Client) PullChan(clientID interface{}) (message chan interface{}, err error) {
	if ch, ok := c.clients[clientID]; ok {
		return ch, nil
	} else {
		err = errors.New("push: no such client")
	}
	return
}
