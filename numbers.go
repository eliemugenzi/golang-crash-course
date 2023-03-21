package main

import "fmt"
func displayNumbers() {
	// Integers
	var age1 int = 20
	var age2 = 30
	age3:= 38

	fmt.Println(age1, age2, age3)

	// Bits & Memory
	var num1 int8 = 21
	var num2 int16 = 215
	var num3 = -128
	var num4 uint = 25
	var num5 uint8 = 255

	var float1 float32 = 25.98 /* Floats */
	var float2 float64 = 333.4424545454
	fmt.Println(num1, num2, num3, num4, num5, float1, float2)
}