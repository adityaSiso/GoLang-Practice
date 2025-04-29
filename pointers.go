package main

import "fmt"

func normalFunction() {
	age := 32
	fmt.Println("Age:", age, "Address:", &age)

	adultYears := getAdultYears(age)
	fmt.Println("Adult Years:", adultYears)
}

func getAdultYears(age int) int {
	fmt.Println("Address inside function:", &age)
	return 18 - age
}

func pointerFunction() {
	var ageAddress *int
	age := 32
	ageAddress = &age

	fmt.Println("Age:", age, "Address:", ageAddress)

	adultYearsbyAddress := getAdultYearsByAddress(ageAddress)
	fmt.Println("Adult Years:", adultYearsbyAddress)
}

func getAdultYearsByAddress(ageAddress *int) int {
	fmt.Println("Address inside function:", ageAddress)
	return 18 - *ageAddress
}
