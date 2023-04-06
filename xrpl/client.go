package xrpl

import (
	"errors"
	"fmt"
	"log"
	"math"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type ClientConfig struct {
	URL                string
	Authorization      string
	Certificate        string
	ConnectionTimeout  time.Duration
	FeeCushion         uint32
	Key                string
	MaxFeeXRP          uint64
	Passphrase         byte
	Proxy              byte
	ProxyAuthorization byte
	Timeout            time.Duration
}

type Client struct {
	config            ClientConfig
	connection        *websocket.Conn
	closed            bool
	mutex             sync.Mutex
	response          *http.Response
	LedgerStream      chan []byte
	ValidationStream  chan []byte
	TransactionStream chan []byte
	PeerStatusStream  chan []byte
	ConsensusStream   chan []byte
	PathFindStream    chan []byte
	err               error
}

func (config *ClientConfig) Validate() error {
	if len(config.URL) == 0 {
		return errors.New("cannot create a new connection with an empty URL")
	}

	if config.ConnectionTimeout < 0 || config.ConnectionTimeout >= math.MaxInt32 {
		return fmt.Errorf("connection timeout out of bounds: %d", config.ConnectionTimeout)
	}

	if config.Timeout < 0 || config.Timeout >= math.MaxInt32 {
		return fmt.Errorf("timeout out of bounds: %d", config.Timeout)
	}

	return nil
}

func NewClient(config ClientConfig) *Client {
	if err := config.Validate(); err != nil {
		panic(err)
	}

	if config.ConnectionTimeout == 0 {
		config.ConnectionTimeout = 60 * time.Second
	}

	client := &Client{
		config:            config,
		LedgerStream:      make(chan []byte, 4),
		ValidationStream:  make(chan []byte, 4),
		TransactionStream: make(chan []byte, 4),
		PeerStatusStream:  make(chan []byte, 4),
		ConsensusStream:   make(chan []byte, 4),
		PathFindStream:    make(chan []byte, 4),
	}
	c, r, err := websocket.DefaultDialer.Dial(config.URL, nil)
	if err != nil {
		client.err = err
		log.Println("Error connecting to xrpl: ", config.URL)
	}
	client.connection = c
	client.response = r
	client.handleResponse()
	fmt.Println("XRPL response: ", r)
	return client
}

func (c *Client) Request(r []byte) error {
	fmt.Println("Sending request: ")
	err := c.connection.WriteMessage(websocket.TextMessage, r)
	if err != nil {
		log.Println("Request error", err)
		return err
	}
	return nil
}

func (c *Client) Subscribe(stream string) error {
	m := fmt.Sprintf("{\"command\":\"subscribe\",\"streams\":[\"%s\"]}", stream)
	err := c.Request([]byte(m))
	if err != nil {
		log.Println("XRPL write error: ", err)
		return err
	}
	return nil
}

func (c *Client) Close() error {
	c.mutex.Lock()
	c.closed = true
	c.mutex.Unlock()

	err := c.connection.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		log.Println("Write close error: ", err)
		return err
	}
	err = c.connection.Close()
	if err != nil {
		log.Println("Write close error: ", err)
		return err
	}
	return nil
}
