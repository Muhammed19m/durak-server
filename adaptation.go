package main

import (
	"errors"
	"fmt"
	durak "serv_durak/logic"
)

func Motion(game *durak.Game, mes Message, id int) (int, error) {
	if mes.Typ_motion == HitCard || mes.Typ_motion == ThrowCard {
		if mes.Card == nil {
			return 0, errors.New("not found card")
		} else if !mes.Card.IsValid() {
			return 0, fmt.Errorf("invalid card %v", *mes.Card)
		}
	}
	switch mes.Typ_motion {
	case Bito:
		return game.Bito(id)
	case Pass:
		return 0, game.Pass(id)
	case Raise:
		return game.Raise(id)
	case HitCard:
		return 0, game.HitCard(id, *mes.Card, mes.Index_battle_cards)
	case ThrowCard:
		return 0, game.ThrowCard(id, *mes.Card)
	default:
		return 0, fmt.Errorf("unknown command %v", mes.Typ_motion)
	}
}
