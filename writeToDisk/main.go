package main

import (
	"fmt"
)

type cards []string

func main() {
	// create a slice
	ls := []string{"card1", "card2"}
	deck := cards(ls)

	fmt.Println(deck)
	// loop over slice
	for _, card := range deck {
		fmt.Println(card)
	}
	var strs string = deck.toString()
	b := []byte(strs)
	fmt.Println(strs, b)
	fmt.Println("Time to write to file!")
	SaveFile(b, "test_file.txt")

	fmt.Println("Time to reload this file in")
	reingested := strFromFile("test_file.txt")
	fmt.Println(reingested)
	ss := BytesToStringList(reingested)
	fmt.Println(ss)
}
