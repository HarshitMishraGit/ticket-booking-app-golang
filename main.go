package main

import "fmt"

func main() {
	// Call the function
	print("Hello, World!")

	var x = 5
	print("variable x = ", x)

	var str string = "Hello, World!"
	fmt.Printf("str = %s\n", str)

	// Read input from the user
	var firstName string
	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Printf("Hello, %s\n", firstName)

	var booking []string
	booking = append(booking, "Hotel")
	fmt.Print("Booking: ", booking, "\n")

	 
}
