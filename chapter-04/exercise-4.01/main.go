package main

import "fmt"

func main() {
	var arr [10]int
	var arr2 [10]int
	fmt.Printf("%#v\n", arr)
	fmt.Println(arr == arr2)
}