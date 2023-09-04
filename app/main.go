package main

import (
	"log"
	"net/http"
)

func main() {
	room := newRoom()

	http.Handle("/ws", room)

	go room.run()

	if err := http.ListenAndServe(":8080", nil) ; err != nil {
		log.Fatal("ListenAndServe : ", err)
	}
}

