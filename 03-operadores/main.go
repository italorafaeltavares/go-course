package main

import "fmt"

func main() {

	//Arithmetic operators
	sub := 80 - 6

	fmt.Println(40 + 7)
	fmt.Println(sub)
	//Operated relationships	fmt.Println(35 > 5)

	age := 25

	fmt.Println(age > 18) // >, >=, ==, <, <=
	// assignment operators
	x := 10
	x += 5 // x = x + 5
	x *= 2
	fmt.Println(x)

	userPass := "1234"
	isAdmin := false
	// if userPass == "1234" && isAdmin == true {
	if userPass == "1234" || isAdmin == true {
		fmt.Println("User logged in")
	} else {
		fmt.Println("Access not permitted")
	}

	a := 10 //1010
	b := 3  //0011
	result := a & b
	fmt.Println(result)

}
