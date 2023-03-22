package main

import (
	"fmt"
	"sort"
	"strings"
)

func displayStrings() {
	// Strings
	var name1 string = "ELie"
	var name2 = "Mugenzi"
	name3:= "Wolton Grant"
	fmt.Println("Hello World")
	fmt.Println(name1, name2, name3)
	name1="Ché"
	name2 = "Traceo"
	fmt.Println(name1, name2, name3)
}

func stringsLib() {

	greeting:= "Hello gophers, I am on my way there!"

	fmt.Println((strings.Contains(greeting, "Hello!")))
	fmt.Println(strings.ReplaceAll(greeting, "Hello", "Hi"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "ph"))

	spliceArr:= strings.Split(greeting, " ")

	fmt.Println(spliceArr, len(spliceArr))

	fmt.Println("The original Value = ", greeting)

	names:= []string {"Che", "Vondelle", "T wizzy", "Avelino"}

	sort.Strings((names))
	fmt.Println(names)
	fmt.Println(sort.SearchStrings(names, "Vondelle"))

	// Loops
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for index, value := range names {
		if index == 1 {
          fmt.Println("Let's skip this one")
		  continue
		}

		if index == 2 {
			fmt.Println("Let's take a break")
			break
		}
		fmt.Println(index, value)
	}

	for _, value := range names {
		fmt.Println(value)
	}

	
}

func cycleNames(n [] string, f func(string)) {
	for _, val:= range n {
		f(val)
	}
}
