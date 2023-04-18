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
	QueueCapacity      int
}

type Client struct {
	config            ClientConfig
	connection        *websocket.Conn
	closed            bool
	mutex             sync.Mutex
	response          *http.Response
	StreamLedger      chan []byte
	StreamTransaction chan []byte
	StreamValidation  chan []byte
	StreamManifest    chan []byte
	StreamPeerStatus  chan []byte
	StreamConsensus   chan []byte
	StreamPathFind    chan []byte
	StreamServer      chan []byte
	StreamDefault     chan []byte
	requestQueue      map[int]string
	nextId            int
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

	if config.QueueCapacity == 0 {
		config.QueueCapacity = 128
	}

	client := &Client{
		config:            config,
		StreamLedger:      make(chan []byte, config.QueueCapacity),
		StreamTransaction: make(chan []byte, config.QueueCapacity),
		StreamValidation:  make(chan []byte, config.QueueCapacity),
		StreamManifest:    make(chan []byte, config.QueueCapacity),
		StreamPeerStatus:  make(chan []byte, config.QueueCapacity),
		StreamConsensus:   make(chan []byte, config.QueueCapacity),
		StreamPathFind:    make(chan []byte, config.QueueCapacity),
		StreamServer:      make(chan []byte, config.QueueCapacity),
		StreamDefault:     make(chan []byte, config.QueueCapacity),
		requestQueue:      make(map[int]string),
		nextId:            0,
	}
	c, r, err := websocket.DefaultDialer.Dial(config.URL, nil)
	if err != nil {
		client.err = err
		return nil
	}
	defer r.Body.Close()
	client.connection = c
	client.response = r
	client.connection.SetPongHandler(client.handlePong)
	go client.handleResponse()
	return client
}

func (c *Client) Ping(message []byte) error {
	if err := c.connection.WriteMessage(websocket.PingMessage, message); err != nil {
		return err
	}
	return nil
}

func (c *Client) NextID() int {
	c.mutex.Lock()
	c.nextId++
	c.mutex.Unlock()
	return c.nextId
}

func (c *Client) Subscribe(stream []byte) error {
	m := fmt.Sprintf(`{"id":"%d","command":"subscribe","streams":["%s"]}`, c.NextID(), stream)
	err := c.Request([]byte(m))
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) Request(req []byte) error {
	fmt.Println("Sending request: ", string(req))
	err := c.connection.WriteMessage(websocket.TextMessage, req)
	if err != nil {
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
