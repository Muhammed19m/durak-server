package game

import "fmt"

type Status struct {
	cards                Cards
	trump                Card
	amount_cards_in_deck int
	battle               []PairCard
}

func (gm *Game) GetStatusPlayer(id int) (Status, error) {
	if i, find := gm.players.SimpleSearchById(id); find {
		ply := gm.players[i]
		return Status{
			*ply.cards,
			gm.trump,
			gm.card_deck.Len(),
			gm.current_battle,
		}, nil
	}
	return Status{}, fmt.Errorf("player id %v not found", id)
}
