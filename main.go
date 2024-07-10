package main

import (
	"fmt"
	"sync"
	"ticket-booking-app/helper"
	"time"
)

var wg = sync.WaitGroup{}

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

		wg.Add(1) // number of go routines to wait for
		go printTicket(firstName)

		if remaingSeats == 0 {
			fmt.Println("All seats are booked")
			break
		}

	}

	wg.Wait() // wait for all go routines to finish
}

func printTicket(firstName string) {
	time.Sleep(5 * time.Second)
	fmt.Println("#############################################")
	fmt.Println("Ticket booked for ", firstName)
	fmt.Println("#############################################")
	wg.Done() // decrement the number of go routines to wait for
}
