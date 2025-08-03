package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Go Tutorial - Map!")

	// Create a map using make
	myMap := make(map[string]int)
	myMap["ONE"] = 1
	myMap["TWO"] = 2
	myMap["THREE"] = 3

	// Print the whole map
	fmt.Println(myMap)				// Output: map[ONE:1 THREE:3 TWO:2]


	// Handle non-existent key
	value, ok := myMap["FOUR"]
	if ok {
		fmt.Println("Exists:", value)
	} else {
		fmt.Println("Key not found")
	}


	// Declare and initialize a map using shorthand syntax
	countries := map[string]string{
		"IN": "India",
		"US": "America",
	}

	// Add a new key-value pair
	countries["NP"] = "Nepal"

	// Delete an element/key
	delete(countries, "US")

	// Print the map
	fmt.Println(countries)			// Output: map[IN:India NP:Nepal US:America]

	// Iterate over the map using for-range
	for key, value := range countries {
		fmt.Printf("%v | %v\n", key, value)
	}
}