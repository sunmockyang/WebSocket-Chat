package Messaging

import (
	"net/http"
)

type Messenger struct {
	connectionManager *ConnectionManager
}

func CreateMessenger() *Messenger {
	return &Messenger{
		connectionManager: CreateConnectionManager(),
	}
}

func (this *Messenger) Start() {
	go this.connectionManager.StartListening()

	http.HandleFunc("/ws", this.HandleRequests)
	err := http.ListenAndServe(":8080", nil)

	checkForError(err, warningGenerator("COULD NOT START HTTP SERVER."), assert)
}

func (this *Messenger) HandleRequests(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	ws, err := upgrader.Upgrade(w, r, nil)
	if !checkForError(err, errorGenerator("Invalid HTTP request. Dropping...")) {
		log(err.Error())
		return
	}

	c := CreateConnection(ws, this.connectionManager)

	this.connectionManager.register <- c
	go c.Listen()
	c.Writer()
}
