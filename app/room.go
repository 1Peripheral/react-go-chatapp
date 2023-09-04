package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Room struct {
	clients map[*Client]bool

	// Join Channel
	joinChan chan *Client

	// Leave Channel
	leaveChan chan *Client

	// Messages sent from clients will be stored on this channel and forwarded to the others.
	forwardChan chan []byte
}

func newRoom() *Room {
	return &Room{
		forwardChan: make(chan []byte),
		joinChan:    make(chan *Client),
		leaveChan:   make(chan *Client),
		clients:     make(map[*Client]bool),
	}
}

func (r *Room) run() {
	for {
		select {
		case client := <-r.joinChan:
			r.clients[client] = true

		case client := <-r.leaveChan:
			delete(r.clients, client)
			close(client.receiveChan)

		case msg := <-r.forwardChan:
			for client := range r.clients {
				client.receiveChan <- msg
			}
		}
	}
}

const (
	socketBufferSize  = 1024
	messageBufferSize = 2561
)

var upgrader = &websocket.Upgrader{
	ReadBufferSize:  socketBufferSize,
	WriteBufferSize: socketBufferSize,
}

func (r *Room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP : ", err)
		return
	}

	client := &Client{
		socket:      socket,
		receiveChan: make(chan []byte, messageBufferSize),
		room:        r,
	}
	r.joinChan <- client
	defer func() { r.leaveChan <- client }()
	go client.write()
	client.read()
}
