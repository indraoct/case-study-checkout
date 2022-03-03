package unit_test

import (
	"case-study-checkout/entity"
	"case-study-checkout/pkg/config"
	checkout2 "case-study-checkout/usecase/checkout"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCase3(t *testing.T){
	t.Run("positive case: Buying more than 3 Alexa Speakers will have a 10% discount on all Alexa speakers", func(t *testing.T) {
		var (
			checkout entity.Checkout
		)
		config := config.Get()

		checkout.Carts = []entity.Cart{
			{
				Sku: "A34",
				Qty: 4,
			},
		}

		dataCheckout,err := checkout2.NewCheckout(config).Checkout(&checkout)
		assert.Equal(t,nil,err)
		assert.Equal(t, 99.98,dataCheckout.Checkout.Total)

	})
}