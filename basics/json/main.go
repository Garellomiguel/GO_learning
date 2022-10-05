package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {
	myJson := `[
		{
			"first_name":"Pepe",
			"last_name": "cruz",
			"hair_color": "blue",
			"has_dog": true
		},
		{
			"first_name":"Pepino",
			"last_name": "cruzado",
			"hair_color": "black",
			"has_dog": false
		}
	]`

	// Load json into a struct
	var unmarshelled []Person //the json came as a slice or list

	err := json.Unmarshal([]byte(myJson), &unmarshelled) // Take a slice of bytes and an iterface
	if err != nil {
		log.Panic("error:", err)
	}

	// log.Printf("unmarshalled: %v", unmarshelled[0].FirstName)
	log.Printf("unmarshalled: %v", unmarshelled)

	// write json from struct
	var mySlice []Person

	var m1 Person
	m1.FirstName = "Wally"
	m1.LastName = "West"
	m1.HairColor = "red"
	m1.HasDog = true

	mySlice = append(mySlice, m1)

	var m2 Person
	m2.FirstName = "Dani"
	m2.LastName = "Pri"
	m2.HairColor = "Black"
	m2.HasDog = false

	mySlice = append(mySlice, m2)

	newJson, err := json.MarshalIndent(mySlice, "", "   ") // same as marshall but with nice sintaxis
	if err != nil {
		log.Panic("error:", err)
	}

	fmt.Println(string(newJson))
}
