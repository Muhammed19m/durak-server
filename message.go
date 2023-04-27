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
	Index_battle_cards int // for HitCard, default 0
	Card               *durak.Card
}

// Bito
// Pass
// Raise
// HitCard
// ThrowCard
