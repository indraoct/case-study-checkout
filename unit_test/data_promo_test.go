package unit_test

import (
	"case-study-checkout/entity"
	"case-study-checkout/pkg/config"
	"case-study-checkout/repository/promo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDataPromoMacbook(t *testing.T){
	t.Run("positive case data promo macbook", func(t *testing.T) {
		var (
			checkout entity.Checkout
		)
		config := config.Get()

		checkout.Carts = []entity.Cart{
			{Sku: "43N23P",
			 Qty: 3,
			},
		}

		Promo,isValidPromo,err := promo.NewPromo(config).PromoMacBook(&checkout)

		assert.Equal(t,nil,err)
		assert.Equal(t,true,isValidPromo)
		assert.NotEqual(t, "", Promo.SKU)

	})
}

func TestDataPromoGoogleHome(t *testing.T){
	t.Run("positive case data promo google home", func(t *testing.T) {
		var (
			checkout entity.Checkout
		)
		config := config.Get()

		checkout.Carts = []entity.Cart{
			{Sku: "120P90",
				Qty: 3,
			},
		}

		Promo,isValidPromo,err := promo.NewPromo(config).PromoGoogleHome(&checkout)

		assert.Equal(t,nil,err)
		assert.Equal(t,true,isValidPromo)
		assert.NotEqual(t, "", Promo.SKU)

	})
}

func TestDataPromoAlexa(t *testing.T){
	t.Run("positive case data promo google home", func(t *testing.T) {
		var (
			checkout entity.Checkout
		)
		config := config.Get()

		checkout.Carts = []entity.Cart{
			{Sku: "A304SD",
				Qty: 4,
			},
		}

		Promo,isValidPromo,err := promo.NewPromo(config).PromoAlexaSpeakers(&checkout)

		assert.Equal(t,nil,err)
		assert.Equal(t,true,isValidPromo)
		assert.NotEqual(t, "", Promo.SKU)

	})
}
