package Messaging

import (
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type connection struct {
	socket  *websocket.Conn
	manager *ConnectionManager
	send    chan []byte
}

func CreateConnection(conn *websocket.Conn, connManager *ConnectionManager) *connection {
	c := connection{
		socket:  conn,
		manager: connManager,
		send:    make(chan []byte),
	}

	c.socket.SetReadLimit(maxMessageSize)

	return &c
}

func (this *connection) Listen() {
	defer func() {
		this.manager.deregister <- this
		this.socket.Close()
	}()
	this.socket.SetReadDeadline(time.Now().Add(pongWait))
	this.socket.SetPongHandler(func(string) error { this.socket.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := this.socket.ReadMessage()
		if !checkForError(err, errorGenerator("Cannot read message...")) {
			break
		}

		this.manager.broadcast <- message
	}
}

func (this *connection) Writer() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		this.socket.Close()
	}()
	for {
		select {
		case message, ok := <-this.send:
			if !ok {
				if !this.write(websocket.CloseMessage, []byte{}) {
					return
				}
			} else {
				if !this.write(websocket.TextMessage, message) {
					return
				}
			}

		case <-ticker.C:
			if !this.write(websocket.PingMessage, []byte{}) {
				return
			}
		}
	}
}

func (this *connection) write(messageType int, data []byte) bool {
	this.socket.SetWriteDeadline(time.Now().Add(writeWait))
	return checkForError(this.socket.WriteMessage(messageType, data), errorGenerator("Failed to write to connection"))
}
