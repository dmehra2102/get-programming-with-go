package main

import "fmt"

func main() {
	userName := "Sir_King_Über"
	runes := []rune(userName)
	for i := 0; i < len(runes); i++ {
		fmt.Print(string(runes[i])," ")
	}
}