package main

import (
	"fmt"
	"ticket-booking-app/helper"
)

func main() {

	var booking []string
	var remaingSeats uint = 5
	for {
		// Read input from the user
		var firstName string
		fmt.Print("Enter your first name: ")
		fmt.Scan(&firstName)
		_ = helper.AddBooking(booking, firstName) // _ is used to decalre a variable that is not used
		fmt.Printf("Hello, %s\n", firstName)
		remaingSeats = remaingSeats - 1
		fmt.Println("Remaining seats : ", remaingSeats)

		if remaingSeats == 0 {
			fmt.Println("All seats are booked")
			break
		}
	}

}
