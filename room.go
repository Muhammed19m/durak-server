package main

import (
	"log"
	durak "serv_durak/logic"
	"time"

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

	s, err := game.GetStatusPlayer(1)
	if err == nil { // everytime err == nil
		room.ply1.Send(&s)
	}
	s, err = game.GetStatusPlayer(2)
	if err == nil { // everytime err == nil
		room.ply2.Send(s)
	}
loop:
	for {
		select {
		case mes := <-room.ply1.Recv:
			_, err := Motion(&game, mes, 1)
			room.ply1.SendText(err.Error())
		case mes := <-room.ply2.Recv:
			_, err := Motion(&game, mes, 2)
			room.ply2.SendText(err.Error())
		case <-time.After(time.Second * 15): // timeout
			room.ply1.SendText("game end, someone didn't fit")
			room.ply2.SendText("game end, someone didn't fit")
			break loop
		}

		s, err = game.GetStatusPlayer(1)
		if err == nil {
			room.ply1.Send(s)
		}
		s, err = game.GetStatusPlayer(2)
		if err == nil {
			room.ply2.Send(s)
		}
	}
	log.Println("Game end")
}
