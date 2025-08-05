package main

import "fmt"

func main() {
	// Main Course
	var total float64 = 2 * 13
	fmt.Println("Sub :", total)

	// Drinks
	total += (4*2.25)
	fmt.Println("Sub :", total)

	// Discount
	total -= 5
	fmt.Println("After Discount :", total)

	// 10% Tip
	tip := total * 0.1
	fmt.Println("Tip :", tip)

	total = total + tip
	fmt.Println("Total:", total)

	// Split bill
	split := total/2
	fmt.Println("Split:",split)

	// Reward every 5th visit
  	visitCount := 24
  	visitCount = visitCount + 1

	remainder := visitCount % 5
	if remainder == 0 {
		fmt.Println("With this visit, you've earned a reward.")
	}
}