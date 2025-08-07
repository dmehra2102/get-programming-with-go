package main

import (
	"encoding/json"
	"os"
	"strings"
)

type ConfigData struct {
	UserName string
	AdditionalProducts []Product
}

var config ConfigData

func LoadConfig() (err error) {
	data,err := os.ReadFile("config.json")
	if err == nil {
		decoder := json.NewDecoder(strings.NewReader(string(data)))
		err = decoder.Decode(&config)
	}
	return
}

func init(){
	err := LoadConfig()
	if err != nil {
		Printfln("Error Loading Config: %v", err.Error())
	}else {
		Printfln("Username: %v", config.UserName)
        Products = append(Products, config.AdditionalProducts...)
	}
}