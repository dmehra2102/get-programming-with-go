package main

import (
	"fmt"
	"time"
)

var (
	Debug       bool
	LogLevel    = "info"
	startUpTime time.Time = time.Now()
)

func main() {
	fmt.Println(Debug,LogLevel,startUpTime)
}