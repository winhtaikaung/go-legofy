package main

import (
	"fmt"

	"./legofy"
)

func main() {
	fmt.Println("Lego My lovely Lego")
	numbers := make([]int, 0)
	numbers = append(numbers, 0)
	numbers = append(numbers, 1)
	numbers = append(numbers, 2)
	numbers = append(numbers, 3)

	fmt.Println(len(numbers))
	fmt.Println(len(legofy.ExtendPalette(numbers, 256, 3)))
}
