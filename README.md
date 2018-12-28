# push
--
    import "github.com/anikhasibul/push"

package push gives you the ability to use push and pull mechanism for
notification or message via websocket or even http client.

## Usage

#### type Client

```go
type Client interface {
	// Push sends a message to all connected clients.
	Push(message interface{})
	// Pull pulls a message for the given clientID.
	Pull(clientID interface{}) (message interface{}, err error)
}
```

Client holds the interface for push and pull mechanism

#### func  NewClient

```go
func NewClient(sessionID, clientID interface{}) Client
```
NewClient returns a client interface.

 `sessionID` means the userID or a groupID.
`clientID` means the deviceID.

A single user (sessionID) can use multiple devices (clientID).
That's why the clientID should be unique for each device.
