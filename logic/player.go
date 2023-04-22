package game

type Player struct {
	cards   *Cards
	progres int
	/*
		progres:
		0 - значит первый ходил
		1 - второй
		и тд...
	*/
	id int
}

func NewPlayer(id int) (p Player) {
	cards := Cards(make([]Card, 0, 6))
	p.cards = &cards
	p.id = id
	return
}

func (p *Player) GiveCard(cards []Card) {
	*p.cards = append(*p.cards, cards...)
}

type Players []Player

func (s Players) Len() int {
	return len(s)
}
func (s Players) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s Players) Less(i, j int) bool {
	return s[i].progres < s[j].progres
}

func (ps Players) SimpleSearchById(srch_id int) (int, bool) {
	for i, p := range ps {
		if p.id == srch_id {
			return i, true
		}
	}
	return 0, false
}
