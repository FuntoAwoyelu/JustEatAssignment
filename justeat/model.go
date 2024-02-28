package justeat

type Restaurant struct {
	Name    string    `json:"name"`
	Address Address   `json:"address"`
	Rating  Rating    `json:"rating"`
	Cuisine []Cuisine `json:"cuisines"`
}

type Address struct {
	Firstline  string `json:"firstLine"`
	City       string `json:"city"`
	PostalCode string `json:"postalCode"`
}

type Rating struct {
	StarRating float64 `json:"starRating"`
}

type Cuisine struct {
	Name string `json:"name"`
}