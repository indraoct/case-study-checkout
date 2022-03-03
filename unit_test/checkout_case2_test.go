package unit_test

import (
	"case-study-checkout/entity"
	"case-study-checkout/pkg/config"
	checkout2 "case-study-checkout/usecase/checkout"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCase2(t *testing.T){
	t.Run("positive case: Buy 3 Google Homes for the price of 2", func(t *testing.T) {
		var (
			checkout entity.Checkout
		)
		config := config.Get()

		checkout.Carts = []entity.Cart{
			{
				Sku: "120P90",
				Qty: 3,
			},
		}

		dataCheckout,err := checkout2.NewCheckout(config).Checkout(&checkout)
		assert.Equal(t,nil,err)
		assert.Equal(t, 99.98,dataCheckout.Total)

	})
}