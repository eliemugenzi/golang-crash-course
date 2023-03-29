package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// Make new bills

func newBill(name string) Bill {
	createdBill := Bill{
		name:  name,
		items: map[string]float64{"pie": 5.99, "cake": 3.99},
		tip:   2.99,
	}

	return createdBill
}

// Format the bill
/* Receiver function*/

func (b Bill) format() string {
	fs := "Bill breakdown: \n"
	var total float64 = 0

	// List items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ... $%v\n", k+":", v)
		total += v
	}

	// Add tip
	fs += fmt.Sprintf("\n%-25v ... $%v\n", "tip:", b.tip)

	// total
	fs += fmt.Sprintf("%-25v ... $%0.2f", "total:", total+b.tip)
	return fs
}

// Update tip
func (b *Bill) updateTip(tip float64) {
	(*b).tip = tip
}

// Add an item to the bill

func (b *Bill) addItem(name string, price float64) {
	(*b).items[name] = price
}

func promptOptions(b Bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose an option(a - Add item, s - Save Bill, t - Add tip): ", reader)

	switch opt {
	case "a":
		itemName, _ := getInput("Item name: ", reader)
		itemPrice, _ := getInput("Item Price: ", reader)
		price, err := strconv.ParseFloat(itemPrice, 64)

		if err != nil {
			fmt.Println("The price must be a number")
			promptOptions(b)
		}

		b.addItem(itemName, price)
		fmt.Println("Items added: ", itemName, price)
		promptOptions(b)
    case "s":
		fmt.Println("Bill saved ---", b.name)
		fmt.Println(b.format())
		b.save()
	case "t":
		tip, _ := getInput("Enter tip amount ($): ", reader)
		t, err := strconv.ParseFloat(tip, 64)

		if err != nil {
			fmt.Println("The tip amount should be a number")
			promptOptions(b)
		}

		b.updateTip(t)
		fmt.Println("Tip added -", tip)
		promptOptions(b)
	default:
		fmt.Println("That was not a valid option... ")
		promptOptions(b)
	}

}

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createBill() Bill {
	reader := bufio.NewReader(os.Stdin)
	name, _ := getInput("Create a bill name: ", reader)

	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)
	return b
}

// Save Bill
func (bill Bill) save(){
   data := [] byte (bill.format())
   err := os.WriteFile("bills/"+bill.name + ".txt", data, 0644)
   if err != nil {
	panic(err)
   }

   fmt.Println("Bill was saved successfully!")
}
