package main

func main() {

	// // cards := []string{"Ace of Spades", newcard()}

	// // Since the deck is declared in deck.go
	// cards := deck{"Ace of Spades", newcard()}

	// cards = append(cards, "Six of spades")

	// // for i, card := range cards {
	// // 	fmt.Println(i, card)
	// // }

	// // Accessing the loop obj func from deck.go
	// cards.print()

	// fmt.Println("Total Card ", cards)

	// NEW CARD FUNC
	// cards := newDeck()
	// fmt.Println(cards)
	// cards.print()

	// value1, value2 := deal(cards, 5)
	// fmt.Println(value1)
	// fmt.Println(value2)
	// fmt.Println("Value 1: ")
	// value1.print()
	// fmt.Println("Value 2 => ")
	// value2.print()
	// greetings := "Hey 	Morning"
	// fmt.Println([]byte(greetings))

	// cards := newDeckFromFile("new_file")
	// cards.print()

	cards := newDeck()
	// cards.print()
	cards.shuffle()
	cards.print()

}

// func newcard() string {
// 	return "Five of Diamonds"
// }
