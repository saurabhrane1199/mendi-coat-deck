package mendicoat_deck

import (
	"math/rand"
	"sort"
	"strings"
)

type Card struct {
	Name, Face, Suit, VerboseName string
	Value                         int
}

type Deck struct {
	Cards []Card
}

func NewDeck(complete bool) Deck {
	if !complete {
		return Deck{Cards: []Card{}}
	}

	var _deck []Card

	for _, s := range suites {
		for v := 0; v < 13; v++ {
			f := ranks[v]
			n := ranks[v] + s
			vn := cards[n]

			c := Card{
				Name:        n,
				Face:        f,
				Suit:        s,
				VerboseName: vn,
				Value:       v + 1,
			}

			_deck = append(_deck, c)

		}
	}

	return Deck{Cards: _deck}
}

func (d *Deck) Shuffle() {
	deck_cards := d.Cards
	rand.Shuffle(len(deck_cards), func(i, j int) {
		deck_cards[i], deck_cards[j] = deck_cards[j], deck_cards[i]
	})
}

func (dest *Deck) DrawFrom(src *Deck, n int) {
	dest.Cards = append(dest.Cards, src.Cards[0:n]...)
	src.Cards = src.Cards[n:]
}

func (d Deck) String() string {
	var g []string
	for _, c := range d.Cards {
		g = append(g, c.VerboseName)
	}
	return strings.Join(g, " ")
}

func (dest *Deck) validateAndRemoveExcessCards(numOfPlayers int) {

	if numOfPlayers < 4 || numOfPlayers%2 == 1 {
		panic("Invalid number of players: must be at least 4 and an even number")
	}

	cards := dest.Cards
	excess := len(cards) % numOfPlayers
	if excess == 0 {
		return
	}

	sort.Slice(cards, func(i, j int) bool {
		return cards[i].Value < cards[j].Value
	})

	dest.Cards = append(cards[:4], cards[4+excess:]...)
	dest.Shuffle()
	return

}

func Deal(src *Deck, numOfPlayers int) []Deck {
	/*
		- Have that many decks as players
		- Shuffle Deck for luck :)
		- Distribute cards equally among the players
			- remove low ranking cards to make the cards equally divisible
			- constraints ( >=4, %2==0)
		- Go in round robin fashion to distribute cards
		- return Cards

	*/

	var hands []Deck
	hands = make([]Deck, numOfPlayers)
	src.validateAndRemoveExcessCards(numOfPlayers)
	src.Shuffle()
	handSize := len(src.Cards) / numOfPlayers

	for i := 0; i < handSize; i++ {
		for j := 0; j < numOfPlayers; j++ {
			hands[j].DrawFrom(src, 1)
		}
	}

	return hands

}

func main() {
	playingDeck := NewDeck(true)
	print(playingDeck.String())

}
