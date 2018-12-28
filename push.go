// package push gives you the ability to use push and pull mechanism for notification or message via websocket or even http client.
package push

import (
	"errors"
)

// Client holds the interface for push and pull mechanism
type Client interface {
	// Push sends a message to all connected clients.
	Push(message interface{})
	// Pull pulls a message for the given clientID.
	Pull(clientID interface{}) (message interface{}, err error)
}

// NewClient returns a client interface.
//	`sessionID` means the userID or a groupID.
//	`clientID` means the deviceID.
//
// A single user (sessionID) can use multiple devices (clientID).
// That's why the clientID should be unique for each device.
func NewClient(sessionID, clientID interface{}) Client {
	ch := make(chan interface{})
	if _, ok := cmap[sessionID]; !ok {
		cmap[sessionID] = &c{
			clients: make(map[interface{}]chan interface{}),
		}
	}
	if clientID != nil {
		cmap[sessionID].clients[clientID] = ch
	}
	return cmap[sessionID]
}

type (
	c struct {
		clients map[interface{}]chan interface{}
	}
	clientmap map[interface{}]*c
)

var cmap = make(clientmap)

func (c *c) Push(message interface{}) {
	for _, v := range c.clients {
		go func(ch chan interface{}) {
			ch <- message
		}(v)
	}
}

func (c *c) Pull(clientID interface{}) (content interface{}, err error) {
	if ch, ok := c.clients[clientID]; ok {
		content = <-ch
	} else {
		err = errors.New("push: no such client")
	}
	return
}
