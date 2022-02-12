package clientmanager

import (
	"fmt"
	"net"
	"sync"
	"time"
)

type Client struct {
	sync.Mutex
	WriteMessage chan []byte
	ReadMessage  chan []byte
	Connection   net.Conn
	ID           string
}

func NewClient(connection net.Conn, id string) *Client {
	client := &Client{
		WriteMessage: make(chan []byte),
		ReadMessage:  make(chan []byte),
		Connection:   connection,
		ID:           id,
	}
	manager.Register(client)
	go client.ReadServe()
	go client.WriteServe()
	return client
}

func (c *Client) Send(message []byte) *Client {
	c.Lock()
	defer c.Unlock()

	c.WriteMessage <- message

	return c
}

func (c *Client) ReadServe() {
	defer func() {
		c.Connection.Close()
		manager.Unregister(c)
	}()

	// SetReadDeadline is if client side does not send any message to server, will send i/o timeout
	// c.Connection.SetReadDeadline(time.Now().Add(10 * time.Second))
	for {
		message := make([]byte, 1024)

		length, err := c.Connection.Read(message)
		if err != nil {
			fmt.Printf("Read error on %s \n", c.ID)
			break
		}
		if length > 0 {
			c.ReadMessage <- message[:length]
		}
	}
}

func (c *Client) WriteServe() {
	ticker := time.NewTicker(5 * time.Second)

	defer func() {
		ticker.Stop()
		c.Connection.Close()
	}()

	for {
		select {
		case message, ok := <-c.WriteMessage:
			c.Connection.SetWriteDeadline(time.Now().Add(10 * time.Second))

			if !ok {
				c.Connection.Write([]byte{})
				return
			}
			c.Connection.Write(message)

		case <-ticker.C:
			c.Connection.SetWriteDeadline(time.Now().Add(10 * time.Second))
			c.Connection.Write([]byte{})
		}
	}
}
