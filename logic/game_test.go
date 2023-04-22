package game

import (
	"testing"
)

func Test2PlayersStaticDeck(t *testing.T) {
	// Колода:
	// 6 ♦, 7 ♦, 8 ♦, 9 ♦, 10 ♦, Валет ♦, Дама ♦, Король ♦, Туз ♦
	// 6 ♥, 7 ♥, 8 ♥, 9 ♥, 10 ♥, Валет ♥, Дама ♥, Король ♥, Туз ♥
	// 6 ♣, 7 ♣, 8 ♣, 9 ♣, 10 ♣, Валет ♣, Дама ♣, Король ♣, Туз ♣
	// 6 ♠, 7 ♠, 8 ♠, 9 ♠, 10 ♠, Валет ♠, Дама ♠, Король ♠, Туз ♠
	// Козырь: 6 ♦
	test_deck := NewCardDeck36x()

	gm := NewGame()
	assert_eq(gm.trump, NewCard(BUBA, 6), t)
	assert_eq(gm.progress, Progress{1, 2}, t)

	assert_eq(gm.states, States{false}, t)
	assert_eq_slice(gm.card_deck.GetSlice(), test_deck[:len(test_deck)-12], t)
	assert_eq_slice(gm.current_battle, []PairCard{}, t)

	assert_eq_slice_card(*gm.players[0].cards, Cards([]Card{{PIKA, 13}, {PIKA, 14}, {PIKA, 9}, {PIKA, 10}, {KRESTI, 14}, {PIKA, 6}}), t)
	assert_eq_slice_card(*gm.players[1].cards, Cards([]Card{{PIKA, 11}, {PIKA, 12}, {PIKA, 7}, {PIKA, 8}, {KRESTI, 12}, {KRESTI, 13}}), t)
	win, err := gm.Bito(1)
	assert_eq(win, 0, t)
	assert_eq(err.Error(), "your turn", t)

	err = gm.Pass(1)
	assert_eq(err.Error(), "your move", t)

	win, err = gm.Raise(1)
	assert_eq(win, 0, t)
	assert_eq(err.Error(), "not your move", t)

	err = gm.ThrowCard(1, NewCard(PIKA, 6))
	if err != nil {
		t.Errorf("%v != %v", err, nil)
	}
	assert_eq_slice_card(*gm.players[0].cards, Cards([]Card{{PIKA, 13}, {PIKA, 14}, {PIKA, 9}, {PIKA, 10}, {KRESTI, 14}}), t)
	err = gm.HitCard(2, NewCard(PIKA, 11), 0)
	if err != nil {
		t.Errorf("%v != %v", err, nil)
	}
	assert_eq_slice_card(*gm.players[1].cards, Cards([]Card{{PIKA, 12}, {PIKA, 7}, {PIKA, 8}, {KRESTI, 12}, {KRESTI, 13}}), t)

	// не дописано
}

func assert_eq[T comparable](a, b T, t *testing.T) {
	if a != b {
		t.Errorf("%v != %v", a, b)
	}
}

func assert_eq_slice[T comparable](a, b []T, t *testing.T) {
	if len(a) != len(b) {
		t.Errorf("len slice1 != len slice2, %v != %v", len(a), len(b))
	} else {
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				t.Errorf("%v != %v", a, b)
				return
			}
		}
	}
}

func assert_eq_slice_card(a, b []Card, t *testing.T) {
	if len(a) != len(b) {
		t.Errorf("len slice1 != len slice2, %v != %v", len(a), len(b))
	} else {
		for i := 0; i < len(a); i++ {
			if a[i].card != b[i].card || a[i].typ != b[i].typ {
				t.Errorf("%v != %v", a, b)
				return
			}
		}
	}
}
