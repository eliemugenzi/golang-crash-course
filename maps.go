package main

import "fmt"

func handleMaps() {
	menu := map[string] float64 {
		"soup": 4.99,
		"salad": 6.99,
		"pie": 2.99,
	}

	fmt.Println(menu)
	fmt.Println(menu["salad"])

	// Looping through maps

	for k, v := range menu {
		fmt.Println(k, ":", v)
	}

	// Ints as key type

	phoneBook := map[int] string {
		250785844487: "Elie",
		2222223333: "Siebel",
	}

	fmt.Println(phoneBook)
	fmt.Println(phoneBook[250785844487])

	menu["salad"] = 9.99
	fmt.Println(menu)
}