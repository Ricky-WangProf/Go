package main

func main() {
	cards := newDeck()

	hand, remainingHands := deal(cards, 5)

	hand.print()
	remainingHands.print()
}
