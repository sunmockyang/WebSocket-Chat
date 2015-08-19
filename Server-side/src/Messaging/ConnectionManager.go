package Messaging

import (
	"fmt"
)

type ConnectionManager struct {
	connections map[*connection]bool
	register    chan *connection
	deregister  chan *connection
	broadcast   chan []byte
}

func CreateConnectionManager() *ConnectionManager {
	return &ConnectionManager{
		connections: make(map[*connection]bool),
		register:    make(chan *connection),
		deregister:  make(chan *connection),
		broadcast:   make(chan []byte),
	}
}

func (this *ConnectionManager) StartListening() {
	for {
		select {
		case c := <-this.register:
			this.RegisterConnection(c)
		case c := <-this.deregister:
			this.DeregisterConnection(c)
		case c := <-this.broadcast:
			this.BroadcastPacket(c)
		}
	}
}

func (this *ConnectionManager) RegisterConnection(conn *connection) {
	log("RegisterConnection")
	this.connections[conn] = true
}

func (this *ConnectionManager) DeregisterConnection(conn *connection) {
	log("DeregisterConnection")

	_, ok := this.connections[conn]

	if ok {
		delete(this.connections, conn)
		close(conn.send)
		conn.socket.Close()
	}
}

func (this *ConnectionManager) BroadcastPacket(message []byte) {
	fmt.Printf("Broadcasting to: %d clients\n", (len(this.connections)))
	fmt.Printf("Received message: %s\n", message)

	for conn := range this.connections {
		select {
		case conn.send <- message:
		default:
			this.deregister <- conn
		}
	}
}
