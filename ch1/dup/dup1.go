package main

import (
	"bufio"
	"fmt"
	"os"
)

func dup1() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	for key, value := range counts{
		if value > 1 {
			fmt.Printf("%s : %d \n", key ,value)
		}
	}

}

func main() {
	// dup1()
	dup2()
}