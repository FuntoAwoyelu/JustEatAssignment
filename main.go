package main

import (
	"log"

	"github.com/FuntoAwoyelu/JustEatAssignment/justeat"
)

func main() {
	// Fetch data from API
	restaurants, err := justeat.RetrieveRestaurant()
	if err != nil {
		log.Fatalf("failed to retrieve restaurants: %v", err)

	}

	// Display first 10 resturants
	justeat.DisplayRestaurants(restaurants)
}
