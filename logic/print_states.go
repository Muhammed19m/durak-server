package game

import "fmt"

func (gm *Game) PrintState() {
	fmt.Print("\nКолода: ")
	gm.card_deck.Inspect(func(c Card) { c.PrintCard(); fmt.Print(", ") })
	fmt.Print("\nТекущая битва: ")
	for _, v := range gm.current_battle {
		v.PrintPairCard()
		fmt.Print(", ")
	}
	fmt.Println()
	for _, ply := range gm.players {
		fmt.Print("Игрок ", ply.id)
		fmt.Println(" Карты: ")
		for _, c := range *ply.cards {
			fmt.Print("\t")
			c.PrintCard()
			fmt.Println()
		}
	}
}

func (gm *Game) PrintDeck() {
	fmt.Print("\nКолода: ")
	gm.card_deck.Inspect(func(c Card) { c.PrintCard(); fmt.Print(", ") })
	fmt.Println()
}

func (gm *Game) PrintBattle() {
	fmt.Println("Текущая битва: ")
	for _, v := range gm.current_battle {
		v.PrintPairCard()
	}
	fmt.Println()
}

func (gm *Game) PrintTrump() {
	fmt.Print("Козырь: ")
	gm.trump.PrintCard()
	fmt.Println()
}

func (gm *Game) PrintPlayerCards() {
	for _, ply := range gm.players {
		fmt.Print("Игрок ", ply.id)
		fmt.Println(" Карты: ")
		for _, c := range *ply.cards {
			fmt.Print("\t")
			c.PrintCard()
			fmt.Println()
		}
	}
}

func (pc *PairCard) PrintPairCard() {
	if pc.hit == nil {
		fmt.Print("\t")
		pc.card.PrintCard()
		fmt.Println(" не отбит")
	} else {
		fmt.Print("\t")
		pc.card.PrintCard()
		fmt.Print(" отбит ")
		pc.hit.PrintCard()
		fmt.Println()
	}
}

func (c *Card) PrintCard() {
	switch {
	case c.card < 11:
		fmt.Print(c.card, " ")
	case c.card == 11:
		fmt.Print("Валет ")
	case c.card == 12:
		fmt.Print("Дама ")
	case c.card == 13:
		fmt.Print("Король ")
	case c.card == 14:
		fmt.Print("Туз ")
	}

	switch c.typ {
	case BUBA:
		fmt.Print("♦")
	case CHERVA:
		fmt.Print("♥")
	case KRESTI:
		fmt.Print("♣")
	case PIKA:
		fmt.Print("♠")
	}
}

func (gm *Game) PrintProgress() {
	fmt.Println("Ходит игрок: ", gm.progress.who)
	fmt.Println("Отбив игрок: ", gm.progress.towhom)
}
