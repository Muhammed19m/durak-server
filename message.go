package main

import durak "serv_durak/logic"

const (
	Bito      = iota
	Pass      = iota
	Raise     = iota
	HitCard   = iota
	ThrowCard = iota
)

type Message struct {
	Typ_motion         int
	Card               *durak.Card
	Index_battle_cards int // for HitCard, default 0
}

// Bito
// Pass
// Raise
// HitCard
// ThrowCard
