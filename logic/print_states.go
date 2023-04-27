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
	if pc.Hit == nil {
		fmt.Print("\t")
		pc.Card.PrintCard()
		fmt.Println(" не отбит")
	} else {
		fmt.Print("\t")
		pc.Card.PrintCard()
		fmt.Print(" отбит ")
		pc.Hit.PrintCard()
		fmt.Println()
	}
}

func (c *Card) PrintCard() {
	switch {
	case c.Card < 11:
		fmt.Print(c.Card, " ")
	case c.Card == 11:
		fmt.Print("Валет ")
	case c.Card == 12:
		fmt.Print("Дама ")
	case c.Card == 13:
		fmt.Print("Король ")
	case c.Card == 14:
		fmt.Print("Туз ")
	}

	switch c.Typ {
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
