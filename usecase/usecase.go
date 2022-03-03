package usecase

import "case-study-checkout/entity"

type ICheckout interface{
	Checkout(checkout *entity.Checkout) (resp entity.CheckoutTotal,err error)
}

type IDiscountRule interface{
	DiscountPrice(checkout *entity.Checkout) (err error)
	DiscountPercentage(checkout *entity.Checkout) (err error)
	DiscountSKUPrice(checkout *entity.Checkout) (err error)
	DiscountSKUPercentage(checkout *entity.Checkout) (err error)
}