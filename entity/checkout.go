package entity

type Checkout struct {
	Carts 	[]Cart `json:"carts"`
	SubTotal float64 `json:"sub_total"`
	Discount float64 `json:"discount"`
	Promos  []Promo  `json:"promos"`
	Total   float64 `json:"total"`
}

type Cart struct {
	Sku 	string	`json:"sku"`
	Qty 	int64	`json:"qty"`
	Price   float64 `json:"price"`
}