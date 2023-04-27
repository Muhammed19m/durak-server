package game

import (
	"errors"
	"fmt"
	"serv_durak/logic/deque"
)

type Game struct {
	card_deck      deque.Deque[Card]
	players        Players
	current_battle []PairCard
	trump          Card
	progress       Progress
	states         States
}

func (gm *Game) GetProgress() Progress {
	return gm.progress
}
func NewGame() Game {
	cards := NewCardDeck36x()
	deq := deque.New[Card]()
	deq.Push_back(cards...)
	game := Game{
		deq,
		[]Player{NewPlayer(1), NewPlayer(2)},
		make([]PairCard, 0, 6),
		cards[0],
		Progress{},
		States{},
	}
	game.startDeal()
	return game
}

func (gm *Game) DealCards() {
	i, _ := gm.players.SimpleSearchById(gm.progress.who)
	need_cards := 6 - len(*gm.players[i].cards)
	if need_cards > 0 {
		*gm.players[i].cards = append(*gm.players[i].cards, gm.card_deck.Pop_back_slice(6-len(*gm.players[i].cards))...)
	}
	for _, p := range gm.players {
		if p.id != gm.progress.towhom {
			need_cards := 6 - len(*p.cards)
			if need_cards > 0 {
				*p.cards = append(*p.cards, gm.card_deck.Pop_back_slice(need_cards)...)
			}
		}
	}
	i, _ = gm.players.SimpleSearchById(gm.progress.towhom)
	need_cards = 6 - len(*gm.players[i].cards)
	if need_cards > 0 {
		*gm.players[i].cards = append(*gm.players[i].cards, gm.card_deck.Pop_back_slice(need_cards)...)
	}
}

func (gm *Game) startDeal() {
	for i := 0; i < 3; i++ {
		for _, p := range gm.players {
			*p.cards = append(*p.cards, gm.card_deck.Pop_back_slice(2)...)
		}
	}
	gm.startProgress()
}

func (gm *Game) ThrowCard(id int, card Card) error {
	if gm.progress.who == id {
		if Contains(gm.current_battle, card) || len(gm.current_battle) == 0 {
			if j, find := gm.players.SimpleSearchById(gm.progress.towhom); find {
				if CanStill(gm.current_battle, &gm.players[j]) > 0 {
					if i, find := gm.players.SimpleSearchById(id); find {
						if card_i, find := gm.players[i].cards.SimpleSearchByCard(card); find {
							gm.players[i].cards.Remove(card_i)
							gm.current_battle = append(gm.current_battle, NewPairCard(card))
							return nil
						}
						return errors.New("no such card")
					}
					return fmt.Errorf("player with id %d does not exist", id)
				}
				return errors.New("no more cards")
			}
			return fmt.Errorf("player with id %d does not exist", gm.progress.towhom)
		}
		return errors.New("this card cannot be thrown")
	}
	return errors.New("not your move")
}

func (gm *Game) HitCard(id int, card Card, index_battle_cards int) error {
	if gm.progress.towhom == id {
		if index_battle_cards < len(gm.current_battle) {
			if i, find := gm.players.SimpleSearchById(id); find {
				if card_i, find := gm.players[i].cards.SimpleSearchByCard(card); find {

					card_who := gm.current_battle[index_battle_cards]
					if card_who.hit == nil {
						if (card_who.card.Card < card.Card && card_who.card.Typ == card.Typ) || (card_who.card.Typ != gm.trump.Typ && card.Typ == gm.trump.Typ) {
							gm.players[i].cards.Remove(card_i)
							gm.current_battle[index_battle_cards].hit = &card
							return nil
						}
						return errors.New("the card is weak to fight back")
					}
					return errors.New("this card has already been canceled")
				}
				return errors.New("no such card")
			}
			return fmt.Errorf("player with id %d does not exist", id)
		}
		return errors.New("no such card exists")
	}
	return errors.New("not your move")
}

func (gm *Game) Raise(id int) (int, error /* > 0 if win, if == 0 then nobody won */) {
	if gm.progress.towhom == id {
		if i, find := gm.players.SimpleSearchById(id); find {
			if gm.states.pass {
				for _, pc := range gm.current_battle {
					if pc.hit != nil {
						*gm.players[i].cards = append(*gm.players[i].cards, pc.card, *pc.hit)
					} else {
						*gm.players[i].cards = append(*gm.players[i].cards, pc.card)
					}
				}
				gm.current_battle = gm.current_battle[:0]
				gm.states.pass = false

				gm.progress.who = gm.progress.towhom
				if gm.players.Len() == gm.progress.who {
					gm.progress.towhom = 1
				} else {
					gm.progress.towhom = gm.progress.who + 1
				}
				gm.progress.who = gm.progress.towhom
				if gm.players.Len() == gm.progress.who {
					gm.progress.towhom = 1
				} else {
					gm.progress.towhom = gm.progress.who + 1
				}
				// 1 2
				// 3 1
				// 1 2 ...

				gm.DealCards()
				for _, ply := range gm.players {
					if len(*ply.cards) == 0 {
						return ply.id, nil
					}
				}
				return 0, nil
			}
			return 0, fmt.Errorf("wait for pass")
		}
		return 0, fmt.Errorf("player with id %d does not exist", id)
	}
	return 0, fmt.Errorf("not your move")
}

func (gm *Game) Pass(id int) error {
	if gm.progress.who == id {
		if len(gm.current_battle) > 0 {
			gm.states.pass = true
			return nil
		}
		return fmt.Errorf("your move")
	}
	return fmt.Errorf("not your move")
}

func (gm *Game) Bito(id int) (int, error /* > 0 if win, if == 0 then nobody won */) {
	if gm.progress.who == id {
		if len(gm.current_battle) == 0 {
			return 0, errors.New("your turn")
		}
		for _, v := range gm.current_battle {
			if v.hit == nil {
				return 0, errors.New("not all cards are beaten")
			}
		}
		gm.current_battle = make([]PairCard, 0)
		gm.DealCards()

		gm.progress.who = gm.progress.towhom
		if gm.players.Len() == gm.progress.who {
			gm.progress.towhom = 1
		} else {
			gm.progress.towhom = gm.progress.who + 1
		}
		for _, ply := range gm.players {
			if len(*ply.cards) == 0 {
				return ply.id, nil
			}
		}
		return 0, nil
	}
	return 0, fmt.Errorf("not your move")
}

func (gm *Game) startProgress() {
	var min_trump_card Card = Card{100, 100}
	id := 1
	for _, ply := range gm.players {
		for _, c := range *ply.cards {

			if c.Typ == gm.trump.Typ && min_trump_card.Card > c.Card {
				min_trump_card = c
				id = ply.id
			}
		}
	}
	gm.progress.who = id
	if gm.players.Len() == id {
		gm.progress.towhom = 1
	} else {
		gm.progress.towhom = id + 1
	}
}

type Progress struct {
	who, towhom int
}

type States struct {
	pass bool
}

func CanStill(sl []PairCard, ply_slugger *Player) int {
	res := 0
	for _, v := range sl {
		if v.hit == nil {
			res++
		}
	}
	return len(*ply_slugger.cards) - res
}
