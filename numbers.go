package main

import (
	"fmt"
	"math"
	"sort"
)
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

func numLib() {
	ages:= []int {1, 22, 443, 11, 60, 44}
	sort.Ints(ages)
	fmt.Println(ages)

	index:= sort.SearchInts(ages, 22)

	fmt.Println(index)
}

func numLoops() {
	x:= 0

	for x < 5 {
		fmt.Println("The Power of X=", x)
		x+=1
	}

	for i := 0; i < 5; i++ {
		fmt.Println("The value of i=", i)
	}
}

type Shape interface {
	getArea() float64
	getCircumference() float64
}

type Square struct {
	length float64
}

type Circle struct {
	radius float64
}

func (square Square) getArea() float64 {
   return math.Pow(square.length, 2)
}

func (circle Circle) getArea() float64 {
  return math.Pi * math.Pow(circle.radius, 2)
}

func (square Square) getCircumference() float64 {
   return square.length * 4
}

func (circle Circle) getCircumference() float64 {
   return 2 * math.Pi * circle.radius
}

func printShapeInfo(shape Shape) {
	fmt.Printf("Area of %T is %0.2f \n", shape, shape.getArea())
	fmt.Printf("The circumference of %T is %0.2f\n", shape, shape.getCircumference())
}