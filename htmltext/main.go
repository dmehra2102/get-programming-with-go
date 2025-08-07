package main

import (
	"html/template"
	"os"
)

func main() {
	t,err := template.ParseFiles("templates/template.html")
	if err == nil {
		t.Execute(os.Stdout, &kayak)
	}else {
		Printfln("Error: %v",err.Error())
	}

}