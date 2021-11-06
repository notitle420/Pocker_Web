package main

import (
	"math/rand"
	"time"

	"github.com/notitle420/pocker_backend_go/check_hand"
)

func main() {
	check_hand.All_card_init()

	a := check_hand.Card{2, 5}
	b := check_hand.Card{4, 8}
	c := check_hand.Card{1, 10}
	d := check_hand.Card{2, 1}
	e := check_hand.Card{2, 10}
	f := check_hand.Card{3, 5}
	g := check_hand.Card{4, 5}

	check_hand.Check_hand(a, b, c, d, e, f, g)
	//fmt.Println("hand", rondom_hand)
	//rondom_hand := make_random_hand()
	//check_hand.Check_hand(rondom_hand[0], rondom_hand[1], rondom_hand[2], rondom_hand[3], rondom_hand[4], rondom_hand[5], rondom_hand[6])

}

func make_random_hand() []check_hand.Card {
	hand := make([]check_hand.Card, 0)
	rand.Seed(time.Now().UnixNano())
	tmp_num := 0
	rand_num := 0
	for i := 0; i < 15; i++ {
		rand_num = rand.Intn(52)
		if rand_num != tmp_num {
			hand = append(hand, check_hand.All_cards[rand_num])
		}
		tmp_num = rand_num
		if len(hand) > 7 {
			break
		}
	}
	return hand
}
