package main

import "log"

func main() {

	cat := "er"

	switch cat {
	case "cat":
		log.Println("IS CAT")

	case "dog":
		log.Println("IS DOG")

	case "fish":
		log.Println("IS FISH")

	default:
		log.Println("IS NOTHING")
	}
}
