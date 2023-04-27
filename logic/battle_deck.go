package game

type PairCard struct {
	Card Card
	Hit  *Card
}

func Contains(cs []PairCard, card Card) bool {
	for _, v := range cs {
		if v.Card.Card == card.Card {
			return true
		} else if v.Hit != nil {
			if v.Hit.Card == card.Card {
				return true
			}
		}
	}
	return false
}

func NewPairCard(card Card) PairCard {
	return PairCard{card, nil}
}
