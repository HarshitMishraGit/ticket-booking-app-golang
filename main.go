package main

import (
	"fmt"
	"ticket-booking-app/helper"
	"time"
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

		go printTicket(firstName)

		if remaingSeats == 0 {
			fmt.Println("All seats are booked")
			break
		}

	}

}

func printTicket(firstName string) {
	time.Sleep(5 * time.Second)
	fmt.Println("#############################################")
	fmt.Println("Ticket booked for ", firstName)
	fmt.Println("#############################################")

}
