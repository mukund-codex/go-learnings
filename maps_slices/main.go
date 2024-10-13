package main

import "log"

type User struct {
	FirstName string
	LastName  string
}

func main() {
	/*
		Maps
	*/
	myMap := make(map[string]User)

	me := User{
		FirstName: "Mukunda",
		LastName:  "Vishwakarma",
	}

	myMap["me"] = me

	log.Println(myMap["me"].FirstName)
	log.Println(myMap["me"].LastName)

	/*
		Slices
	*/
}
