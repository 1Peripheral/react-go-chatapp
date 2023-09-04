package main

import (
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	id string

	socket *websocket.Conn

	receiveChan chan []byte

	room *Room
}

func (c *Client) read() {
	defer c.socket.Close()
	for {
		_, msg, err := c.socket.ReadMessage()
		if err != nil {
			log.Println()
			return
		}
		c.room.forwardChan <- msg

	}
}

func (c *Client) write() {
	defer c.socket.Close()
	for msg := range c.receiveChan {
		err := c.socket.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			return
		}
	}
}
