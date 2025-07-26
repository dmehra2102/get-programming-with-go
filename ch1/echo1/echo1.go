package main

import (
	"fmt"
	"os"
)

func main() {
	sep, s := "", ""
	for _, value := range os.Args[1:] {
		s += sep + value
		sep = " "
	}

	fmt.Println(os.Args[0])
	fmt.Println(s)
	exercise1_2()
}

func exercise1_2(){
	for index,value := range os.Args[1:] {
		fmt.Println("Index :",index , " Value: ",value)
	}
}