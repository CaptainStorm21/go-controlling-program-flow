package main

import "fmt"

func main() {

	u1 := User{
		ID:        1,
		FirstName: "Albert",
		LastName:  "Moran",
	}
	u2 := User{
		ID:        3,
		FirstName: "Josephine",
		LastName:  "Garden",
	}

	if u1.ID == u2.ID {
		fmt.println("Same user")
	} else if u1.FirstName == u2.FirstName {
		fmt.println("they dont have a same first name")
	} else {
		fmt.println("Not the same user")
	}

	if u1.FirstName != u2.LastName {
		fmt.println("not the same student")
	}
}