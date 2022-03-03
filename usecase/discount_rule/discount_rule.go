package discount_rule

import (
	"case-study-checkout/entity"
	"case-study-checkout/pkg/config"
	"case-study-checkout/usecase"
)

type ObjDiscountRule struct {
	Config 	config.Configuration
}

func NewDicountRule(config config.Configuration) usecase.IDiscountRule{
	return &ObjDiscountRule{Config: config}
}

func (d *ObjDiscountRule) DiscountPrice(checkout *entity.Checkout) (err error){

	return nil
}

func (d *ObjDiscountRule) DiscountPercentage(checkout *entity.Checkout) (err error){

	return nil
}

func (d *ObjDiscountRule) DiscountSKUPrice(checkout *entity.Checkout) (err error){

	return nil
}

func (d *ObjDiscountRule) DiscountSKUPercentage(checkout *entity.Checkout) (err error){

	return nil
}