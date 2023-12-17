package model

type House struct {
	ID          string  `bson:"_id,omitempty"`
	Address     string  `bson:"address"`
	Description string  `bson:"description"`
	City        string  `bson:"city"`
	State       string  `bson:"state"`
	PostalCode  int     `bson:"postalCode"`
	Price       float64 `bson:"price"`
	ForSale     bool    `bson:"forSale"`
	Available   bool    `bson:"available"` //Shows whether the house is available for rent or not
}
