/*
* This program is free software; you can redistribute it and/or modify
* it under the terms of the GNU General Public License as published by
* the Free Software Foundation; either version 2 of the License, or
* (at your option) any later version.
*
* This program is distributed in the hope that it will be useful,
* but WITHOUT ANY WARRANTY; without even the implied warranty of
* MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
* GNU General Public License for more details.
*
* You should have received a copy of the GNU General Public License
* along with this program; if not, see <http://www.gnu.org/licenses/>.
*
* Copyright (C) Hasibul Hasan (Anik), 2018
 */

// Package push gives you the ability to use push and pull mechanism for notification or message via websocket or even http client.
package push

import (
	"errors"
	"sync"
)

type (
	clients   map[interface{}]ClientChan
	clientmap map[interface{}]*Session
)

var mu sync.Mutex
var cmap = make(clientmap)

// Session holds the methods for push and pull mechanism
type Session struct {

	//	`MaxChannelBuffer` means the maximum buffered message on a client channel. `make(chan interface{},MaxChannelBuffer)`.
	// Default value is 10
	MaxChannelBuffer int
	id               interface{}
	clients          clients
}

// ClientChan holds a chan of interface type to provide type flexibility on pushed message
type ClientChan chan interface{}

// NewSession returns a client session.
//
// A single user (sessionID) can use multiple devices (clientID).
// That's why the clientID should be unique for each device/client/connection.
//
//
//	`sessionID` means the userID or a groupID. Once a `Session` receives a message, it pushes the message to all registered client for this session.
//
func NewSession(sessionID interface{}) *Session {
	// if the session doesn't exist, create a new one
	if _, ok := cmap[sessionID]; !ok {
		mu.Lock()
		cmap[sessionID] = &Session{
			id:               sessionID,
			MaxChannelBuffer: 10,
			clients:          make(clients),
		}
		mu.Unlock()
	}
	return cmap[sessionID]
}

// NewClient returns a `Client`.
//
// It creates a new client if the given `clientID` does not exist.
//
// It returns the existing client for the given `clientID` if it already exists.
//
// it panics if the given `clientID` is nil.
//
//
// A single user (sessionID) can use multiple devices (clientID).
// That's why the clientID should be unique for each device/client/connection.
func (s *Session) NewClient(clientID interface{}) *Client {

	// if the client exists, return the existing one
	if _, ok := cmap[s.id].
		clients[clientID]; ok {
		return &Client{
			clientID: clientID,
			session:  cmap[s.id],
		}
	}
	// panic for nil
	if clientID == nil {
		panic("push: clientID is nil")
	}
	// create a new one
	clientChan := make(ClientChan, s.MaxChannelBuffer)
	mu.Lock()
	cmap[s.id].clients[clientID] = clientChan
	mu.Unlock()

	// return the created client
	return &Client{
		clientID: clientID,
		session:  cmap[s.id],
	}
}

// Push sends a message to all connected clients on the given session.
/*
	s := push.NewClient(userID, nil,0)
	s.Push("Hello world!")
*/
func (s *Session) Push(message interface{}) {
	wg := sync.WaitGroup{}
	for key := range s.clients {
		wg.Add(1)
		go func(k interface{}) {
			defer wg.Done()
			s.clients[k] <- message
		}(key)
	}
	wg.Wait()
}

// DeleteClient deletes the given client from the current session.
// It's safe to delete a non-existent client.
func (s *Session) DeleteClient(clientID interface{}) {
	delete(s.clients, clientID)
}

// Len returns the length/count of active clients on current session.
func (s *Session) Len() int {
	return len(s.clients)
}

// Clients returns the keys/IDs/names of active clients on current session.
func (s *Session) Clients() []interface{} {
	ret := []interface{}{}
	for k := range s.clients {
		ret = append(ret, k)
	}
	return ret
}

// DeleteSelf deletes the current session from memory.
func (s *Session) DeleteSelf() {
	delete(cmap, s.id)
}

// pull ...
func (s *Session) pull(clientID interface{}) (content interface{}, err error) {
	if ch, ok := s.clients[clientID]; ok {
		content = <-ch
	} else {
		err = errors.New("push: no such client")
	}
	return
}

// pullChan ...
func (s *Session) pullChan(clientID interface{}) (message ClientChan, err error) {

	if ch, ok := s.clients[clientID]; ok {
		return ch, nil
	}
	err = errors.New("push: no such client")

	return
}

// DeleteSession deletes the given session from memory.
// It's safe to delete a non-existent session.
func DeleteSession(sessionID interface{}) {
	delete(cmap, sessionID)
}
