package main

import (
	"fmt"
)

// func main() {
// 	colors := map[string]string{
// 		"red":   "#ff0000",
// 		"green": "#00FF00",
// 	}

// 	var numbers map[string]string

// 	fmt.Println(colors)
// 	fmt.Println(numbers)

// 	friends := make(map[string]string)

// 	friends["Rajesh"] = "Gurgaon"
// 	fmt.Println(friends)

// 	delete(friends, "Rajesh")
// 	fmt.Println(friends)

// }

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00FF00",
		"white": "#ffffff",
	}
	// fmt.Println(colors)
	printMap(colors)

}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for ", color, "is ", hex)
	}
}
