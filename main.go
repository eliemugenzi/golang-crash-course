package main

import "fmt"


func main() {
	displayStrings()
	displayNumbers()
	formatStrings()
	handleArrays()
	stringsLib()
	numLib()
	numLoops()
	greetAFriend("Peshyensi")
	fact:= factorial(6)
	fmt.Println("Fact=", fact)

	cycleNames([] string {"Che", "Headie", "Michael", "RV", "Dushane"}, greetAFriend)

	radius := 1.09
	area := getCircleArea(radius)
	fmt.Println("The area of the circle=", area)
	fmt.Printf("Circle area=%0.3f\n", area)

	value0, value1 := getInitials("Ché Wolton-Grant")
	value2, value3 := getInitials("Elie")

	fmt.Println(value0, value1)
	fmt.Println(value2, value3)

	handleMaps()
	handlePointers()	
	myBill := newBill("Elie's Bill")
	myBill.updateTip(20)
	myBill.addItem("biscuit", 4.99)
	myBill.addItem("Capuccino", 2.99)
	fmt.Println(myBill.format())

	userBill := createBill()
	fmt.Println(userBill)

	promptOptions(userBill)

	shapes := [] Shape {
		Square { length: 15.2 },
		Circle {radius: 20},
		Circle {radius: 20.8},
		Square{length: 19.7},
	}

	for _, v := range shapes {
		printShapeInfo(v)
		fmt.Println("----")
	}
}
