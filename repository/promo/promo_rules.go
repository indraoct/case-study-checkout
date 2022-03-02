package promo

import (
	"case-study-checkout/constants"
	"case-study-checkout/entity"
	"case-study-checkout/pkg/config"
	"case-study-checkout/repository"
	"case-study-checkout/repository/products"
)

type ObjPromo struct {
	Config config.Configuration
}

func NewPromo(config config.Configuration) repository.IPromo{
	return &ObjPromo{Config: config}
}

func (p *ObjPromo) PromoMacBook(checkout entity.Checkout) (promo entity.Promo, isPromoValid bool,err error){
	isPromoValid = false
	promo.MinimumQty = 1
	var qty int64 = 0
	for _, cart := range checkout.Carts{
		if cart.Sku == "43N23P" && cart.Qty >= promo.MinimumQty {
			isPromoValid = true
			qty = cart.Qty
			break
		}
	}

	if isPromoValid == true{
		promo.SKU = "43N23P"
		promo.PromoType = constants.PROMO_TYPE_DISCOUNT_SKU_PRICE
		productPromo,_ :=products.NewProducts(p.Config).GetProductsBySKU("234234")
		promo.PromoSKU = []entity.SKU{
			{
				Sku: "234234",
				Discount: productPromo[0].Price*float64(qty),
			},
		}
	}

	return promo,isPromoValid, nil
}

func (p *ObjPromo) PromoGoogleHome(checkout entity.Checkout) (promo entity.Promo, isPromoValid bool,err error){
	isPromoValid = false
	promo.MinimumQty = 2
	for _, cart := range checkout.Carts{
		if cart.Sku == "120P90" && cart.Qty >= promo.MinimumQty {
			isPromoValid = true
			break
		}
	}

	if isPromoValid == true {
		promo.SKU = "120P90"
		promo.PromoType = constants.PROMO_TYPE_DISCOUNT_PRICE
		productPromo,_ :=products.NewProducts(p.Config).GetProductsBySKU("43N23P")
		promo.DiscountPrice = productPromo[0].Price

	}

	return promo,isPromoValid, nil
}

func (p *ObjPromo) PromoAlexaSpeakers(checkout entity.Checkout) (promo entity.Promo, isPromoValid bool,err error){

	isPromoValid = false
	promo.MinimumQty = 4
	var qty int64 = 0
	for _, cart := range checkout.Carts{
		if cart.Sku == "A304SD" && cart.Qty >= promo.MinimumQty {
			isPromoValid = true
			qty = cart.Qty
			break
		}
	}

	if isPromoValid == true {
		promo.SKU = "A304SD"
		promo.PromoType = constants.PROMO_TYPE_DISCOUNT_SKU_PERCENTAGE
		productPromo,_ :=products.NewProducts(p.Config).GetProductsBySKU("A304SD")
		promo.PromoSKU = []entity.SKU{
			{
				Sku: "A304SD",
				Discount: (0.01 * productPromo[0].Price) * float64(qty),
			},
		}
	}
	return promo,isPromoValid, nil
}