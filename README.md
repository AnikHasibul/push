# push
--
    import "github.com/anikhasibul/push"

Package push gives you the ability to use push and pull mechanism for
notification or message via websocket or even http client.

## Usage

#### func  DeleteSession

```go
func DeleteSession(sessionID interface{})
```
DeleteSession deletes the given session from memory. It's safe to delete a
non-existent session.

#### type ClientChan

```go
type ClientChan chan interface{}
```

ClientChan holds a chan of interface type to provide type flexibility on pushed
message

#### type Session

```go
type Session struct {
}
```

Session holds the methods for push and pull mechanism

#### func  NewSession

```go
func NewSession(sessionID, clientID interface{}, maxChannelBuffer int) *Session
```
NewSession returns a client session.

A single user (sessionID) can use multiple devices (clientID). That's why the
clientID should be unique for each device/client/connection.

    `sessionID` means the userID or a groupID. Once a `Session` receives a message, it pushes the message to all registered client for this session.

    `clientID` means the deviceID. A new client will be created on the given session if the `sessionID` is not `nil`.

    `maxChannelBuffer` means the maximum buffered message on a client channel. `make(chan interface{},maxChannelBuffer)`

#### func (*Session) Clients

```go
func (s *Session) Clients() []interface{}
```
Clients returns the keys/IDs/names of active clients on current session.

#### func (*Session) DeleteClient

```go
func (s *Session) DeleteClient(clientID interface{})
```
DeleteClient deletes the given client from the current session. It's safe to
delete a non-existent client.

#### func (*Session) Len

```go
func (s *Session) Len() int
```
Len returns the length/count of active clients on current session.

#### func (*Session) Pull

```go
func (s *Session) Pull(clientID interface{}) (content interface{}, err error)
```
Pull pulls a message for the given clientID.

    s := push.NewClient(userID, clID,100)
    defer s.DeleteClient(clID)
    for {
    	msg, err := s.Pull(clID):
    	if err != nil {
    		panic(err)
    	}
    	fmt.Println(msg)
    }

#### func (*Session) PullChan

```go
func (s *Session) PullChan(clientID interface{}) (message ClientChan, err error)
```
PullChan returns a channel for receiving messages for the given clientID

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

Exclusively usable with websockets

#### func (*Session) Push

```go
func (s *Session) Push(message interface{})
```
Push sends a message to all connected clients on the given session.

    s := push.NewClient(userID, nil,0)
    s.Push("Hello world!")
