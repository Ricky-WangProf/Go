package main

import (
	"fmt"
)

func main() {
	evenAndOddSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, value := range evenAndOddSlice {
		if value%2 == 0 {
			fmt.Println(evenAndOddSlice[i], "is even.")
		} else {
			fmt.Println(evenAndOddSlice[i], "is odd.")
		}
	}
}
