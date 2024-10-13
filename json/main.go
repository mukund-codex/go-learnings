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

	myJson := `
[
    {
        "first_name": "Mukunda",
        "last_name": "Vishwakarma",
        "hair_color": "black",
        "has_dog": false
    },
    {
        "first_name": "Sanjana",
        "last_name": "Priya",
        "hair_color": "brown",
        "has_dog": true
    }
]`

	var unmarshalled []Person
	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println("Error unmarshalling json", err)
	}

	log.Printf("unmarshalled: %v", unmarshalled)

	//write json from struct

	var mySlice []Person

	var m1 Person
	m1.FirstName = "Kabir"
	m1.LastName = "Deewan"
	m1.HairColor = "blonde"
	m1.HasDog = false

	mySlice = append(mySlice, m1)

	var m2 Person
	m2.FirstName = "Anish"
	m2.LastName = "Pandey"
	m2.HairColor = "black"
	m2.HasDog = false

	mySlice = append(mySlice, m2)

	newJson, err := json.MarshalIndent(mySlice, "", "     ")
	if err != nil {
		log.Println("Error Marshalling", err)
	}

	fmt.Print(string(newJson))
}
