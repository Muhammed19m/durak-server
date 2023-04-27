package game

import "fmt"

type Status struct {
	Cards                Cards
	Trump                Card
	Amount_cards_in_deck int
	Battle               []PairCard
	Progress             Progress
}

func (gm *Game) GetStatusPlayer(id int) (Status, error) {
	if i, find := gm.players.SimpleSearchById(id); find {
		ply := gm.players[i]
		return Status{
			*ply.cards,
			gm.trump,
			gm.card_deck.Len(),
			gm.current_battle,
			gm.progress,
		}, nil
	}
	return Status{}, fmt.Errorf("player id %v not found", id)
}
