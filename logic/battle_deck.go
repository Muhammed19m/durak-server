package game

type PairCard struct {
	card Card
	hit  *Card
}

func Contains(cs []PairCard, card Card) bool {
	for _, v := range cs {
		if v.card.Card == card.Card {
			return true
		} else if v.hit != nil {
			if v.hit.Card == card.Card {
				return true
			}
		}
	}
	return false
}

func NewPairCard(card Card) PairCard {
	return PairCard{card, nil}
}
