package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize, WriteBufferSize: socketBufferSize}

type Room struct {
	// Forward is a channel that holds incoming messages to be forwarded to
	// other clients.
	forward chan []byte
	// join is a channel for clients wishing to join the room.
	join chan *Client
	// leave is a channel for clients wishing to leave the room.
	leave   chan *Client
	clients map[*Client]bool
}

// Instantiate Room.
func newRoom() *Room {
	return &Room{
		forward: make(chan []byte),
		join:    make(chan *Client),
		leave:   make(chan *Client),
		clients: make(map[*Client]bool),
	}
}

// Message broker method.
func (r *Room) run() {
	for {
		select {
		// Join Room
		case client := <-r.join:
			r.clients[client] = true
		// Leave the room
		case client := <-r.leave:
			delete(r.clients, client)
			close(client.send)
		// Forward msg to all clients
		case msg := <-r.forward:
			for client := range r.clients {
				client.send <- msg
			}
		}
	}
}

// Instantiate client and read/write on HTTP handler.
func (r *Room) ServeHTTP(w http.ResponseWriter, rq *http.Request) {
	socket, err := upgrader.Upgrade(w, rq, nil)
	if err != nil {
		log.Fatal("Room ServeHTTP: ", err)
		return
	}

	client := &Client{
		socket: socket,
		send:   make(chan []byte, messageBufferSize),
		room:   r,
	}
	r.join <- client
	defer func() { r.leave <- client }()
	// Write messages from r.forward (c.send) to socket
	go client.Write()
	// Writemessages from c.socket to c.r.forward
	client.Read()
}
