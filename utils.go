package mendicoat_deck

import (
	"sort"
	"strings"
)

const rankOrder = "2345678910JQKA"

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

func isCurrentCardHigher(currentCard Card, leadingCard Card) bool {

	if leadingCard.Face == "" {
		return true
	}

	if strings.Index(rankOrder, currentCard.Face) < strings.Index(rankOrder, leadingCard.Face) {
		return false
	} else {
		return true
	}

}
