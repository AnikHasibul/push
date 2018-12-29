# push
--
    import "github.com/anikhasibul/push"

package push gives you the ability to use push and pull mechanism for
notification or message via websocket or even http client.

## Usage

#### type Client

```go
type Client struct {
}
```

Client holds the interface for push and pull mechanism

#### func  NewClient

```go
func NewClient(sessionID, clientID interface{}) *Client
```
NewClient returns a client interface.

    `sessionID` means the userID or a groupID.
    `clientID` means the deviceID.

A single user (sessionID) can use multiple devices (clientID). That's why the
clientID should be unique for each device.

#### func (*Client) Pull

```go
func (c *Client) Pull(clientID interface{}) (content interface{}, err error)
```
Pull pulls a message for the given clientID.

    c := push.NewClient(userID, clID)
    for {
    	msg, err := c.Pull(clID):
    	if err != nil {
    		panic(err)
    	}
    	fmt.Println(msg)
    }

#### func (*Client) PullChan

```go
func (c *Client) PullChan(clientID interface{}) (message chan interface{}, err error)
```
PullChan returns a channel for receiving messages for the given clientID

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

Extremely usable with websockets

#### func (*Client) Push

```go
func (c *Client) Push(message interface{})
```
Push sends a message to all connected clients.

    c := push.NewClient(userID, nil)
    c.Push("Hello world!")
