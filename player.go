package main

import (
	"log"

	"github.com/gorilla/websocket"
)

type Player struct {
	conn *websocket.Conn
	Recv chan Message
}

func NewPlayer(conn *websocket.Conn) *Player {
	ply := &Player{conn, make(chan Message, 20)}
	go ply.recv()
	return ply
}

func (p *Player) Send(mes any) error {
	return p.conn.WriteJSON(mes)
}

func (p *Player) SendText(mes string) error {
	return p.conn.WriteMessage(websocket.TextMessage, []byte(mes))
}

func (p *Player) recv() {
	for {
		mes := Message{}
		err := p.conn.ReadJSON(mes)
		if err == nil {
			p.Recv <- mes
		} else {
			log.Println("invalid message ", mes, " err: ", err)
		}
	}
}
