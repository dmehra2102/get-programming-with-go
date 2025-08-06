package main

import "fmt"

func main() {
	arr := [...]string{
		"Ready ",
		"Get ",
		"Go ",
		"to ",
	}

	fmt.Print(arr[1],arr[0],arr[3],arr[2])
}