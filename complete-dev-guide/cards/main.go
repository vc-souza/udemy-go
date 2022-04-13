package main

func main() {
	// cards := newDeck()

	// cards.print()

	// hand, remainingCards := deal(cards, 5)

	// fmt.Println("Hand")
	// fmt.Println(len(hand))
	// hand.saveToFile("hand.txt")

	// fmt.Println("Cards Remaining:")
	// fmt.Println(len(remainingCards))
	// remainingCards.saveToFile("cards.txt")

	cards := newDeck()

	cards.shuffle()
	cards.print()
}
