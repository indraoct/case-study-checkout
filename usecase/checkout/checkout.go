package checkout

import (
	"case-study-checkout/entity"
	"case-study-checkout/pkg/config"
	"case-study-checkout/usecase"
)

type ObjCheckout struct {
	Config 	config.Configuration
}

func NewCheckout(config config.Configuration) usecase.ICheckout{
	return &ObjCheckout{Config: config}
}

func (c *ObjCheckout) Checkout(checkout *entity.Checkout) (resp entity.CheckoutTotal,err error){

	return resp,nil
}