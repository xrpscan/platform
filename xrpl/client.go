package xrpl

import (
	"errors"
	"fmt"
	"log"
	"math"
	"net/http"
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
	config       ClientConfig
	connection   *websocket.Conn
	response     *http.Response
	LedgerStream chan LedgerStream
	err          error
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

	client := &Client{config: config}
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

func (c *Client) handleResponse() error {
	go func() {
		for {
			var v interface{}
			err := c.connection.ReadJSON(&v)
			if err != nil {
				log.Println("XRPL read error: ", err)
			}
			fmt.Println(v)
		}
	}()
	return nil
}

func (c *Client) IsConnected() bool {
	return c.connection != nil
}

func (c *Client) Close() error {
	c.connection.Close()
	return nil
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
