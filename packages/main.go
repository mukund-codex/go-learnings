package main

import (
	"log"

	"github.com/mukund-codex/myniceprogram/helpers"
)

func main() {

	var myVar helpers.SomeType
	myVar.TypeName = "some name"

	log.Println(myVar.TypeName)

}
