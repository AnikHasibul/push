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
	clients clients
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
//	`clientID` means the deviceID. A new client will be created on the given session if the `sessionID` is not `nil`.
//
//	`maxChannelBuffer` means the maximum buffered message on a client channel. `make(chan interface{},maxChannelBuffer)`
//
func NewSession(sessionID, clientID interface{}, maxChannelBuffer int) *Session {
	// if the session doesn't exist, create a new one
	if _, ok := cmap[sessionID]; !ok {
		mu.Lock()
		cmap[sessionID] = &Session{
			clients: make(clients),
		}
		mu.Unlock()
	}
	// if the clientexists, return the existing one
	if _, ok := cmap[sessionID].clients[clientID]; ok {
		return cmap[sessionID]
	}
	// create a new client
	if clientID != nil {
		clientChan := make(ClientChan, maxChannelBuffer)
		mu.Lock()
		cmap[sessionID].clients[clientID] = clientChan
		mu.Unlock()
	}
	// return the created session
	return cmap[sessionID]
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

// Pull pulls a message for the given clientID.
/*
	s := push.NewClient(userID, clID,100)
	defer s.DeleteClient(clID)
	for {
		msg, err := s.Pull(clID):
		if err != nil {
			panic(err)
		}
		fmt.Println(msg)
	}
*/
func (s *Session) Pull(clientID interface{}) (content interface{}, err error) {
	if ch, ok := s.clients[clientID]; ok {
		content = <-ch
	} else {
		err = errors.New("push: no such client")
	}
	return
}

// PullChan returns a channel for receiving messages for the given clientID
/*
	s := push.NewClient(userID, clID,100)
	defer s.DeleteClient(clID)
	ch, err := s.PullChan(clID):
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
// Exclusively usable with websockets
func (s *Session) PullChan(clientID interface{}) (message ClientChan, err error) {

	if ch, ok := s.clients[clientID]; ok {
		return ch, nil
	}
	err = errors.New("push: no such client")

	return
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

// DeleteSession deletes the given session from memory.
// It's safe to delete a non-existent session.
func DeleteSession(sessionID interface{}) {
	delete(cmap, sessionID)
}
