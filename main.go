package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", ":8080", "http service address")

var upgrader = websocket.Upgrader{
	ReadBufferSize:  128,
	WriteBufferSize: 128,
}

func main() {
	flag.Parse()

	serv := NewServer()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		log.Println("User connected")
		conn, err := upgrader.Upgrade(w, r, nil)
		if err == nil {
			log.Println("conn upgrade")
			serv.Flow_players <- conn
		} else {
			log.Println("error upgrade")
		}
	})

	log.Println("Server started")
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
