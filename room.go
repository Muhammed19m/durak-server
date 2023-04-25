package main

import (
	durak "serv_durak/logic"

	"github.com/gorilla/websocket"
)

type Room struct {
	ply1, ply2 *Player
}

func NewRoom(con1, con2 *websocket.Conn) Room {
	ply1 := NewPlayer(con1)
	ply2 := NewPlayer(con2)
	return Room{ply1, ply2}
}

func (room Room) Run() {
	game := durak.NewGame()
	game.DealCards()

	// players_id := map[*Player]int{room.ply1: 1, room.ply2: 2}

	for {
		select {
		case mes := <-room.ply1.Recv:
			_, err := Motion(&game, mes, 1)
			room.ply1.SendText(err.Error())
		case mes := <-room.ply2.Recv:
			_, err := Motion(&game, mes, 2)
			room.ply2.SendText(err.Error())
		}
	}
}