package main

import "fmt"

func formatStrings() {
	// Printing & Formatting Strings

	name:= "Elie"
	age:= 26

	fmt.Print("Hello, ")
	fmt.Print("World! \n")
	fmt.Print("New Line\n")
	fmt.Println("Hello, Croydons!")
	fmt.Println("My Name is", name, "and I am", age, "years old!")

	// Printf(Formatted Strings) %_ = format specifier
	fmt.Printf("My name is %v and I am %v years old!\n", name, age)
	fmt.Printf("My name is %q and I am %v years old!\n", name, age)

}
