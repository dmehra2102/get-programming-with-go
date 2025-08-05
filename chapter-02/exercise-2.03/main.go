package main

import (
	"fmt"
	"time"
)

func main() {
	dayBorn := time.Saturday

	switch dayBorn {
	case time.Monday:
		fmt.Println("Monday's child is fair of face")
	case time.Tuesday:
		fmt.Println("Tuesday's child is full of grace")
	case time.Wednesday:
		fmt.Println("Wednesday's child is full of woe")
	case time.Thursday:
		fmt.Println("Thursday's child has far to go")
	case time.Friday:
		fmt.Println("Friday's child is loving and giving")
	case time.Saturday:
		fmt.Println("Saturday's child works hard for a living")
	case time.Sunday:
		fmt.Println("Sunday's child is bonny and blithe")
	default:
		fmt.Println("Error, day born not valid")
	}

	switch dayBorn {
	case time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday:
		fmt.Println("Born on a weekday")
	case time.Saturday, time.Sunday:
		fmt.Println("Born on the weekend")
	default:
		fmt.Println("Error, day born not valid")
	}

	switch dayBorn := time.Sunday;{
	case dayBorn == time.Saturday || dayBorn == time.Sunday:
		fmt.Println("Born on the weekend")
	default:
		fmt.Println("Born some other day")
	}
}
