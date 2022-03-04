package checkout

import (
	"case-study-checkout/entity"
	"case-study-checkout/pkg/config"
	"case-study-checkout/pkg/helper"
	"case-study-checkout/repository/products"
	"case-study-checkout/usecase"
	"case-study-checkout/usecase/discount_rule"
)

type ObjCheckout struct {
	Config 	config.Configuration
}

func NewCheckout(config config.Configuration) usecase.ICheckout{
	return &ObjCheckout{Config: config}
}

func (c *ObjCheckout) Checkout(checkout *entity.Checkout) (resp entity.Checkout,err error){
	//get data price
	i := 0
	for _,cart := range checkout.Carts{
		dataProduct,_ :=products.NewProducts(c.Config).GetProductsBySKU(cart.Sku)
		cart.Price = helper.FloatRound(float64(cart.Qty)*dataProduct[0].Price,2)
		checkout.Carts[i].Price = cart.Price
		checkout.Carts[i].Name = dataProduct[0].Name
		checkout.SubTotal += cart.Price
		checkout.Total += cart.Price
		i++
	}


	//discount rules
	discount_rule.NewDicountRule(c.Config).DiscountPrice(checkout)
	discount_rule.NewDicountRule(c.Config).DiscountPercentage(checkout)
	discount_rule.NewDicountRule(c.Config).DiscountSKUPrice(checkout)
	discount_rule.NewDicountRule(c.Config).DiscountSKUPercentage(checkout)

	resp = *checkout

	return resp,nil
}