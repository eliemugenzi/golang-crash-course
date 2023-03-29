package main
import "fmt"

func updateName(x *string) {
	*x = "Wolton-Grant"
}

func handlePointers () {
	// Pointers
	name := "Elie"
	fmt.Println("The memory address of name:", &name)

	m := &name
	fmt.Println("Memory address:", m)
	fmt.Println("Value at memory address:", *m)

	updateName(m)

	fmt.Println("The updated name: ", name)
}