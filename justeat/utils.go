package justeat

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func RetrieveRestaurant() ([]Restaurant, error) {
	var postCode string
	fmt.Println("Please enter a postcode")
	fmt.Scan(&postCode)

	resp, err := http.Get(fmt.Sprintf("https://uk.api.just-eat.io/discovery/uk/restaurants/enriched/bypostcode/%s", postCode))
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve data from API: %v", err)
	}
	defer resp.Body.Close()

	var data struct {
		Restaurants []Restaurant `json:"restaurants"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("failed to decode JSON: %v", err)
	}

	return data.Restaurants, nil
}

func DisplayRestaurants(restaurants []Restaurant) {
	for i, v := range restaurants {
		if i >= 10 {
			break
		}
		fmt.Printf("Restaurant: %d\n", i+1)
		fmt.Printf("Restaurant Name: %s\n", v.Name)
		fmt.Printf("Restaurant Address: %s, %s, %s\n", v.Address.Firstline, v.Address.City, v.Address.PostalCode)
		fmt.Printf("Restaurant Cuisines: %s\n", FormatCuisines(v.Cuisine))
		fmt.Printf("Restaurant Rating: %.1f\n", v.Rating.StarRating)
		fmt.Println("-----------------------------------------------------------------")
	}
}
func FormatCuisines(cuisines []Cuisine) string {
	var listCuisines []string
	for _, v := range cuisines {
		listCuisines = append(listCuisines, v.Name)
	}
	return strings.Join(listCuisines, ", ")
}
