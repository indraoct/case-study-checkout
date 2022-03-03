package unit_test

import (
	"case-study-checkout/pkg/config"
	"case-study-checkout/repository/products"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAllProducts(t *testing.T){
	t.Run("positive case get all products", func(t *testing.T) {
		config := config.Get()
		products,err := products.NewProducts(config).GetAllProducts()

		assert.Equal(t,nil,err)
		assert.NotEqual(t, 0, len(products))

	})
}

func TestGetProductsBySKU(t *testing.T){
	t.Run("positive case get products by SKU", func(t *testing.T) {
		config := config.Get()
		products,err := products.NewProducts(config).GetProductsBySKU("234234")

		assert.Equal(t,nil,err)
		assert.NotEqual(t, 0, len(products))

	})
}
