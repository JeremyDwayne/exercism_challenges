package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	ruleset := map[string]int{
		"ace": 11, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7,
		"eight": 8, "nine": 9, "ten": 10, "jack": 10, "queen": 10, "king": 10,
	}

	val, ok := ruleset[card]
	if ok {
		return val
	}
	return 0
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	hand := ParseCard(card1) + ParseCard(card2)
	dealer := ParseCard(dealerCard)
	switch {
	case hand == 22:
		return "P"
	case hand == 21:
		if dealer < 10 {
			return "W"
		}
		return "S"
	case hand >= 17 && hand <= 20:
		return "S"
	case hand >= 12 && hand <= 16:
		if dealer >= 7 {
			return "H"
		}
		return "S"
	default:
		return "H"
	}
}
