package unit_test

import (
	"case-study-checkout/entity"
	"case-study-checkout/pkg/config"
	checkout2 "case-study-checkout/usecase/checkout"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCase1(t *testing.T){
	t.Run("positive case: Each sale of a MacBook Pro comes with a free Raspberry Pi B", func(t *testing.T) {
		var (
			checkout entity.Checkout
		)
		config := config.Get()

		checkout.Carts = []entity.Cart{
			{
				Sku: "43N23P",
				Qty: 1,
			},
			{
				Sku: "234234",
				Qty: 1,
			},
		}

		dataCheckout,err := checkout2.NewCheckout(config).Checkout(&checkout)
		assert.Equal(t,nil,err)
		assert.Equal(t, 5399.99,dataCheckout.Checkout.Total)

	})
}
