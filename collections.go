// Arrays, Slices, and Maps
package main

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

func collection_main() {
	// Array
	var productName [4]string
	// Slice
	prices := []float64{10.99, 45.99, 99.99, 4.99}
	fmt.Println(prices[:4])
	fmt.Println(productName)
	for i := 0; i < len(prices); i++ {
		fmt.Println(prices[i])
	}
	slicedLeft := prices[1:]
	slicesRight := slicedLeft[:2]
	fmt.Println("Sliced Left", slicedLeft, "Lenght", len(slicedLeft), "Capacity", cap(slicedLeft))
	fmt.Println("Sliced Right", slicesRight, "Lenght", len(slicesRight), "Capacity", cap(slicesRight))

	prices = append(prices, 105.99)
}

func practice() {
	// Array
	hobbies := []string{"Football", "Gardening", "Riding"}
	fmt.Println(hobbies)

	make_ = make([]string, 2, 3)

	slicedHobbies := hobbies[0:2]
	reSliced := slicedHobbies[1:3]
	// Unpacking example
	hobbies = append(hobbies, slicedHobbies...)
	fmt.Println(reSliced)

	for index, value := range hobbies {
		fmt.Println(index, value)
	}
}

func map_practice() {
	websites := map[string]string{
		"Google": "www.google.com",
		"Amazon": "www.amazon.com",
	}
	make_ = make(map[string]int64, 3)

	fmt.Println(websites)
	fmt.Println(websites["Meta"])
	websites["Meta"] = "www.meta.com"
	delete(websites, "Google")
	fmt.Println(websites)

	for key, value := range websites {
		fmt.Println(key, value)
	}
}
