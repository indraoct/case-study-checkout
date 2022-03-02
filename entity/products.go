package entity

import "case-study-checkout/pkg/config"

type Products struct {
	Sku 	string					`json:"sku"`
	Name    string					`json:"name"`
	Price 	float64 				`json:"price"`
	Qty 	int64 					`json:"qty"`
	Config  config.Configuration	`json:"-"`
}
