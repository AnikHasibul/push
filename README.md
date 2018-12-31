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

#### func  Exists

```go
func Exists(sessionID interface{}) bool
```
Exists returns true if the given session exists.

#### type Client

```go
type Client struct {
}
```

Client holds the methods for a perticular client.

#### func (*Client) Close

```go
func (c *Client) Close()
```
Close closes a client channel/connection

#### func (*Client) DeleteSelf

```go
func (c *Client) DeleteSelf()
```
DeleteSelf deletes the current client from the current session.

#### func (*Client) Key

```go
func (c *Client) Key() interface{}
```
Key returns the current clientID/name/key.

#### func (*Client) KeyString

```go
func (c *Client) KeyString() string
```
KeyString returns the current clientID/name/key in string type.

#### func (*Client) Pull

```go
func (c *Client) Pull() (content interface{}, err error)
```
Pull pulls a message for the current client.

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

#### func (*Client) PullChan

```go
func (c *Client) PullChan() (ClientChan, error)
```
PullChan returns a channel for receiving messages for the current client.

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

Exclusively usable with websockets

#### type ClientChan

```go
type ClientChan chan interface{}
```

ClientChan holds a chan of interface type to provide type flexibility on pushed
message

#### type Session

```go
type Session struct {

	//	`MaxChannelBuffer` means the maximum buffered message on a client channel. `make(chan interface{},MaxChannelBuffer)`.
	// Default value is 10
	MaxChannelBuffer int
}
```

Session holds the methods for push and pull mechanism

#### func  NewSession

```go
func NewSession(sessionID interface{}) *Session
```
NewSession returns a client session.

A single user (sessionID) can use multiple devices (clientID). That's why the
clientID should be unique for each device/client/connection.

    `sessionID` means the userID or a groupID. Once a `Session` receives a message, it pushes the message to all registered client for this session.

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

#### func (*Session) DeleteSelf

```go
func (s *Session) DeleteSelf()
```
DeleteSelf deletes the current session from memory.

#### func (*Session) Exists

```go
func (s *Session) Exists(clientID interface{}) bool
```
Exists returns true if the given client exists.

#### func (*Session) Len

```go
func (s *Session) Len() int
```
Len returns the length/count of active clients on current session.

#### func (*Session) NewClient

```go
func (s *Session) NewClient(clientID interface{}) *Client
```
NewClient returns a `Client`.

It creates a new client if the given `clientID` does not exist.

It returns the existing client for the given `clientID` if it already exists.

it panics if the given `clientID` is nil.

A single user (sessionID) can use multiple devices (clientID). That's why the
clientID should be unique for each device/client/connection.

#### func (*Session) Push

```go
func (s *Session) Push(message interface{})
```
Push sends a message to all connected clients on the given session.

    s := push.NewClient(userID, nil,0)
    s.Push("Hello world!")
