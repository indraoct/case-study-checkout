package discount_rule

import (
	"case-study-checkout/entity"
	"case-study-checkout/pkg/config"
	"case-study-checkout/pkg/helper"
	"case-study-checkout/repository/promo"
	"case-study-checkout/usecase"
)

type ObjDiscountRule struct {
	Config 	config.Configuration
}

func NewDicountRule(config config.Configuration) usecase.IDiscountRule{
	return &ObjDiscountRule{Config: config}
}

func (d *ObjDiscountRule) DiscountPrice(checkout *entity.Checkout) (err error){
	dataPromo,isValidPromo, err :=promo.NewPromo(d.Config).PromoGoogleHome(checkout)

	if err != nil{
		return err
	}

	if isValidPromo == true{
		for _,cart := range checkout.Carts{
			for _, promoSku := range dataPromo.PromoSKU{
				if cart.Sku == promoSku.Sku{
					checkout.Total = helper.FloatRound(checkout.Total - promoSku.Discount,2)
					checkout.Promos = append(checkout.Promos,dataPromo)
					checkout.Discount += promoSku.Discount
				}
			}
		}
	}

	return nil
}

func (d *ObjDiscountRule) DiscountPercentage(checkout *entity.Checkout) (err error){
	//TODO:
	return nil
}

func (d *ObjDiscountRule) DiscountSKUPrice(checkout *entity.Checkout) (err error){
	dataPromo,isValidPromo, err :=promo.NewPromo(d.Config).PromoMacBook(checkout)

	if err != nil{
		return err
	}

	if isValidPromo == true{
		for _,cart := range checkout.Carts{
			for _, promoSku := range dataPromo.PromoSKU{
				if cart.Sku == promoSku.Sku{
					checkout.Total = helper.FloatRound(checkout.Total - promoSku.Discount,2)
					checkout.Promos = append(checkout.Promos,dataPromo)
					checkout.Discount += promoSku.Discount
				}
			}
		}
	}

	return nil
}

func (d *ObjDiscountRule) DiscountSKUPercentage(checkout *entity.Checkout) (err error){
	dataPromo,isValidPromo, err :=promo.NewPromo(d.Config).PromoAlexaSpeakers(checkout)

	if err != nil{
		return err
	}

	if isValidPromo == true{
		for _,cart := range checkout.Carts{
			for _, promoSku := range dataPromo.PromoSKU{
				if cart.Sku == promoSku.Sku{
					checkout.Total = helper.FloatRound(checkout.Total - promoSku.Discount,2)
					checkout.Promos = append(checkout.Promos,dataPromo)
					checkout.Discount += promoSku.Discount
				}
			}
		}
	}

	return nil
}