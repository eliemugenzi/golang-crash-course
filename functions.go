package main

import (
	"fmt"
	"math"
	"strings"
)

func greetAFriend(name string) {
	fmt.Printf("Hello my good friend %v\n", name)
}

func factorial(num int) int {
   if num >= 1 {
	return factorial(num - 1) * num
   } else {
	return 1
   }

}

func getCircleArea(radius float64) float64 {
	return math.Pi * math.Pow(radius, 2)
}

// Return Multiple values

func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")
	var initials [] string
	for _, v := range names {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"

}
