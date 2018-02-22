package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#green",
		"blue":  "#blue",
	}

	// colors := make(map[string]string)

	// colors["red"] = "#ff0000"

	// delete(colors, "red")

	fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color + " is " + hex)
	}
}
