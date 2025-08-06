package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	Day := time.Now().Weekday()
	Hour := time.Now().Hour()

	fmt.Println("Day: ", Day, "Hour: ", Hour)

	appName := "HTTPCHECKER"
	action := "BASIC"
	date := time.Now()

	LogFileName := appName + "_" + action + "_" + strconv.Itoa(date.Year()) + "_" + date.Month().String() + "_" + strconv.Itoa(date.Day()) + ".log"

	fmt.Println("The name of the logfile is: ", LogFileName)
}
