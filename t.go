package main

import "fmt"

func main() {
	var userName string
	var newArray [50]string
	newArray2 := []string{}
	fmt.Println("Please enter your name")
	fmt.Scan(&userName)
	fmt.Printf("Your name is %s \n", userName)
	var bookings = [50]string{"Nana", "mehmet", "ece"}
	fmt.Println(bookings[1])
	newArray[0] = userName
	fmt.Printf("for 1: %v ", newArray[0])
	newArray2 = append(newArray[:], userName)
	fmt.Printf("for 2: %v ", newArray2[:])

	fmt.Printf("%T", newArray2[:])
}
