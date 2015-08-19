package main

import (
	"Messaging"
)

func main() {
	startMessage()
	server := Messaging.CreateMessenger()
	server.Start()
}

func startMessage() {
	println(" _______  __   __  _______  _______ ")
	println("| W.S.  ||  | |  ||   _   ||       |")
	println("|     __||  |_|  ||  |_|  ||_     _|")
	println("|    |__ |   _   ||   _   |  |   |  ")
	println("|_______||__| |__||__| |__|  |___|  ")
	println("______________________Sunmock-2015__")
}
