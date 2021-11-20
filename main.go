package main

// var deckSize int
// variable can be initialize outside function but can be assign value only inside the function
func main() {
	// cards := newCard() Initialize assign variable in single line , type that 'newCard()' function return will assign as type of variable.
	// cards := deck{"Ace of Diamonds", newCard()}

	// cards := newDeck()
	// cards.print()

	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()

	// cards := newDeck()  // For creating new deck of cards
	// cards.saveToFile("my_cards") //Save decks of card to file named called "my_cards"

	cards := newDeckFromFile("my_cards")

	cards.print()
}

/////////////////////////////////////////////////////////
// package main

// import "fmt"

// type book string

// func (b book) printTitle() {
//     fmt.Println(b)
// }

// func main() {
//     var b book = "Harry Potter"
//     b.printTitle()
// }

// type laptopSize float64

// func (this laptopSize) getSizeOfLaptop() laptopSize {
//     return this
// }
