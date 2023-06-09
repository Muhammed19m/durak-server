package game

const (
	_      = iota
	BUBA   = iota
	CHERVA = iota
	KRESTI = iota
	PIKA   = iota
)

type Card struct {
	Typ  int
	Card int
}

type Cards []Card

func (cs *Cards) Remove(i int) {
	after := (*cs)[i+1:]
	*cs = append((*cs)[:i], after...)
}

func (cs *Cards) SimpleSearchByCard(card Card) (int, bool) {
	for i, c := range *cs {
		if c == card {
			return i, true
		}
	}
	return 0, false
}

func NewCard(tp, val int) (card Card) {
	if tp > PIKA && tp < 1 && val < 6 && val > 14 {
		return NewCard(BUBA, 6) // this will be changed in the future
	}
	card.Card = val
	card.Typ = tp
	return
}

func (card *Card) IsValid() bool {
	if card.Typ > PIKA && card.Typ < 1 && card.Card < 6 && card.Card > 14 {
		return false
	}
	return true
}

func NewCardDeck36x() []Card {
	deck := []Card{
		NewCard(BUBA, 6),
		NewCard(BUBA, 7),
		NewCard(BUBA, 8),
		NewCard(BUBA, 9),
		NewCard(BUBA, 10),
		NewCard(BUBA, 11),
		NewCard(BUBA, 12),
		NewCard(BUBA, 13),
		NewCard(BUBA, 14),

		NewCard(CHERVA, 6),
		NewCard(CHERVA, 7),
		NewCard(CHERVA, 8),
		NewCard(CHERVA, 9),
		NewCard(CHERVA, 10),
		NewCard(CHERVA, 11),
		NewCard(CHERVA, 12),
		NewCard(CHERVA, 13),
		NewCard(CHERVA, 14),

		NewCard(KRESTI, 6),
		NewCard(KRESTI, 7),
		NewCard(KRESTI, 8),
		NewCard(KRESTI, 9),
		NewCard(KRESTI, 10),
		NewCard(KRESTI, 11),
		NewCard(KRESTI, 12),
		NewCard(KRESTI, 13),
		NewCard(KRESTI, 14),

		NewCard(PIKA, 6),
		NewCard(PIKA, 7),
		NewCard(PIKA, 8),
		NewCard(PIKA, 9),
		NewCard(PIKA, 10),
		NewCard(PIKA, 11),
		NewCard(PIKA, 12),
		NewCard(PIKA, 13),
		NewCard(PIKA, 14),
	}
	return deck
}
