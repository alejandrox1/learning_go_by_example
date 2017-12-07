package main

import (
	"github.com/gorilla/websocket"
)

// A chatting user.
type Client struct {
	// Web socket for this client.
	socket *websocket.Conn
	// Channel on which messages are sent.
	send chan []byte
	// Room in which the client is chatting.
	room *Room
}

// Read messages from the socket and forward them to the Room's forward
// channel.
func (c *Client) Read() {
	defer c.socket.Close()
	for {
		_, msg, err := c.socket.ReadMessage()
		if err != nil {
			return
		}
		// Send messages to Room.forward channel.
		c.room.forward <- msg
	}
}

// Write messages from the Client's send channel to the socket.
func (c *Client) Write() {
	defer c.socket.Close()
	// Read messages from Client.send channel.
	for msg := range c.send {
		err := c.socket.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			return
		}
	}
}
