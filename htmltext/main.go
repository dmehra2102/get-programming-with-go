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

	allTemplates,err := template.ParseFiles("templates/template.html","templates/extras.html")
	if err == nil {
		allTemplates.ExecuteTemplate(os.Stdout, "extras.html", &kayak)
	}else {
		Printfln("Error: %v",err.Error())
	}

	allTemplate,err := template.ParseGlob("templates/*.html")
	if err == nil {
		for _,t := range allTemplate.Templates() {
			Printfln("Template name: %v", t.Name())
		}
	}else {
		Printfln("Error : %v",err.Error())
	}
}