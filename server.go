package main

import (
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type Server struct {
	flow_players chan *websocket.Conn
	mx           sync.Mutex
	players      []*websocket.Conn
}

func (s *Server) Listen() {
	for {
		ply := <-s.flow_players
		s.mx.Lock()
		s.players = append(s.players, ply)
		s.mx.Unlock()
	}
}

func (s *Server) StartGame() {
	for {
		s.mx.Lock()
		if len(s.players) > 2 {
			ply1 := s.players[len(s.players)-1]
			ply2 := s.players[len(s.players)-2]
			s.players = s.players[:len(s.players)-1]
			room := NewRoom(ply1, ply2)
			go room.Run()
		}
		s.mx.Unlock()
		<-time.NewTicker(time.Millisecond * 100).C
	}
}
