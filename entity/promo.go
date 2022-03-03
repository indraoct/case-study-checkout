package entity

type Promo struct {
	SKU  					string	`json:"sku"`
	MinimumQty 				int64 	`json:"minimum_qty"`
	PromoType 				int 	`json:"promo_type"`
	Discount   		        float64 `json:"discount"`
	Percentage 				float64	`json:"percentage"`
	PromoSKU 				[]SKU   `json:"promo_sku"`
}

type SKU struct {
	Sku 	  string 	`json:"sku"`
	Discount  float64	`json:"discount"`
	Percentage float64  `json:"percentage"`
}