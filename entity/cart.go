package entity

type Checkout struct {
	Carts 	[]Cart `json:"carts"`
	SubTotal float64 `json:"sub_total"`
	Discount float64 `json:"discount"`
	Total   float64 `json:"total"`
}

type Cart struct {
	Sku 	string	`json:"sku"`
	Qty 	int64	`json:"qty"`
}