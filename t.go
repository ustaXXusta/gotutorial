package main

import "fmt"

func main() {
	var userName string
	fmt.Println("Please enter your name")
	fmt.Scan(&userName)
	fmt.Printf("Your name is %s", userName)

}
