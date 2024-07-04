package main

// import "fmt"

/*
function with return and basic print statements

*/
// func main() {

// 	//Mannual initialization
// 	// var card string = "Ace of Spades"

// 	//Dynamic datatype initialization
// 	card:= "Ace of Spades"

// 	card = "Five of Diamonds"

// 	fmt.Println(card)

// 	card=newCard()
// 	fmt.Println(card)

// }

// func newCard () string {
// 	return "Five of spade"
// }

/*
Arrays
*/

// func main(){
// 	cards:= []string{"Ace","Bijoy"};

// 	cards = append(cards, "Six")

// 	fmt.Println(cards);

// 	//Print each value
// 	for i,card := range cards {
// 		fmt.Println(i, card)
// 	}
// }

/*
Custom type declarations

*/

// func main() {
// 	cards := deck{"Ace", "Bijoy"}

// 	cards = append(cards, "Six")

// 	fmt.Println(cards)

// 	//Print each value
// 	// for i, card := range cards {
// 	// 	fmt.Println(i, card)
// 	// }

// 	cards.print()
// }

// Execute using
//go run hello.go deck.go

/*
cards func call to create

*/

// func main() {
// 	cards := newDeck()

// 	fmt.Println(cards)

// 	cards.print()
// }

/*
sub arrays
*/

// func main(){
// 	cards:=newDeck()
// 	hand, remaining_cards:=deal(cards,5)
// 	hand.print()
// 	remaining_cards.print()

// }

/*
Convert to byte slice

*/

// func main() {
// 	greeting := "Hi there!"
// 	fmt.Println([]byte(greeting))
// }

/*
Joining array of strings
*/

// func main() {
// 	card := newDeck()

// 	fmt.Println(card.toString())
// }

/*
Save file to harddrive
*/

// func main(){
// 	cards:=newDeck()
// 	cards.saveToFile("mycards")
// }

/*
read from file
*/
// func main() {
// 	cards := newDeckFromFile("mycards")
// 	cards.print()
// }

/*
randomize a deck of cards
*/

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
