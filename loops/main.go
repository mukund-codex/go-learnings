package main

import "log"

func main() {

	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}

	var users []User
	users = append(users, User{"Sanjana", "Priya", "sanjana@gmail.com", 25})
	users = append(users, User{"Mukunda", "Vishwakarma", "mukunda@gmail.com", 28})
	users = append(users, User{"Kabir", "Deewan", "kabir@gmail.com", 27})
	users = append(users, User{"Anish", "Pandey", "anish@gmail.com", 23})

	for _, l := range users {
		log.Println(l.FirstName, l.LastName, l.Email, l.Age)
	}
}
