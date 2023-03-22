package main

import "fmt"

func handleArrays() {

	// Arrays
	var ages [3]int = [3]int {23,34,56}
	var ages2 = [4]int {12,23,434, 55}
	fmt.Println(ages, len(ages))
	fmt.Println(ages2, len(ages2))

	ages[2] = 100

	names:= [6]string {"Elie", "Che", "Bryant", "Kaninho", "Big Zuu", "Re-Mark"}

	fmt.Println(names, len(names))

	// Slices(use the arrays under the hood!)

	var scores = []int {100, 200, 500, 5000}
	scores = append(scores, 30)

	fmt.Println(scores, len(scores))

	// Slice ranges
	rangeOne:= names[1:4]
	rangeTwo:= names[2:]
	rangeThree:= names[:3]

	rangeOne = append(rangeOne, "Headie One")

	fmt.Println(rangeOne)
	fmt.Println(rangeTwo)
	fmt.Println(rangeThree)

}