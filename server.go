package main

import (
	"log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type Server struct {
	Flow_players chan *websocket.Conn
	mx           sync.Mutex
	players      []*websocket.Conn
}

func NewServer() *Server {
	serv := &Server{
		make(chan *websocket.Conn, 128),
		sync.Mutex{},
		make([]*websocket.Conn, 0, 10),
	}
	go serv.listen()
	go serv.runGame()
	return serv
}

func (s *Server) listen() {
	for {
		ply := <-s.Flow_players
		s.mx.Lock()
		s.players = append(s.players, ply)
		s.mx.Unlock()
	}
}

func (s *Server) runGame() {
	for {
		s.mx.Lock()
		if len(s.players) >= 2 {
			ply2 := s.players[len(s.players)-1]
			ply1 := s.players[len(s.players)-2]
			s.players = s.players[:len(s.players)-2]
			room := NewRoom(ply1, ply2)
			log.Println("Room run")
			go room.Run()
		}
		s.mx.Unlock()
		<-time.NewTicker(time.Millisecond * 100).C
	}
}
