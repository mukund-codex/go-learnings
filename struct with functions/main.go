package main

import "log"

type myStruct struct {
	FirstName string
}

/*
	all values of myStruct is assigned to var m,
	we can directly create functions on struct like this and can get values.
*/
func (m *myStruct) printFirstName() string {
	return m.FirstName
}

func main() {
	var myVar myStruct
	myVar.FirstName = "Mukunda"

	myVar2 := myStruct{
		FirstName: "Sanjana",
	}

	log.Println("myVar is set to ", myVar.printFirstName())
	log.Println("myVar2 is set to ", myVar2.printFirstName())
}
