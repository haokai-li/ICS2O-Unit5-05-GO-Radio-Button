// Created by: haokai
// Created on: May 2021
// This program displays, "Radio-Button"

package main

import (
	"fmt"
)

func main() {

	// This function does addition
	var age int
	var day string

	// input
	fmt.Println("This program is about weekend and age .")
	fmt.Println()
	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)
	fmt.Print("Enter the day of the week: ")
	fmt.Scanln(&day)

	// detect
	if (day == "saturday" || day == "sunday") && (age >= 18 || age < 18) {
		// relax
		fmt.Println("Time to relax for the weekend!")
	} else if !(day == "tuesday" || day == "thursday") && age < 18 {
		// go to school
		fmt.Println("Time for school!")
	} else if !(day == "tuesday" || day == "thursday") && age >= 18 {
		// go to work
		fmt.Println("Time to go to work!")
	}

	fmt.Println("\nDone.")
}
