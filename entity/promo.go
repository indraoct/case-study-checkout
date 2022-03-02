package entity

type Promo struct {
	SKU  					string	`json:"sku"`
	MinimumQty 				int64 	`json:"minimum_qty"`
	PromoType 				int 	`json:"promo_type"`
	DiscountPercentage 		float64 `json:"discount_percentage"`
	DiscountPrice    		float64 `json:"discount_price"`
	PromoSKU 				[]SKU   `json:"promo_sku"`
}

type SKU struct {
	Sku 	  string 	`json:"sku"`
	Discount  float64	`json:"discount"`
}