package main

import "log"

type myStruct struct {
	FirstName string
}

func main() {
	var myVar myStruct
	myVar.FirstName = "Mukunda"

	myVar2 := myStruct{
		FirstName: "Sanjana",
	}

	log.Println("myVar is set to ", myVar.FirstName)
	log.Println("myVar2 is set to ", myVar2.FirstName)
}
