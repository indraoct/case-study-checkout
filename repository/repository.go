package repository

import "case-study-checkout/entity"

type IProducts interface {
	GetAllProducts() (arrProducts []entity.Products,err error)
	GetProductsBySKU(sku string) (products []entity.Products,err error)
}

type IPromo interface {
	PromoMacBook(checkout entity.Checkout) (promo entity.Promo, isPromoValid bool,err error)
	PromoGoogleHome(checkout entity.Checkout) (promo entity.Promo, isPromoValid bool,err error)
	PromoAlexaSpeakers(checkout entity.Checkout) (promo entity.Promo, isPromoValid bool,err error)
}
