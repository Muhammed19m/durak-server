package game

type PairCard struct {
	card Card
	hit  *Card
}

func Contains(cs []PairCard, card Card) bool {
	for _, v := range cs {
		if v.card.card == card.card {
			return true
		} else if v.hit != nil {
			if v.hit.card == card.card {
				return true
			}
		}
	}
	return false
}

func NewPairCard(card Card) PairCard {
	return PairCard{card, nil}
}
