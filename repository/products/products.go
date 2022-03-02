package products

import (
	"case-study-checkout/entity"
	"case-study-checkout/pkg/config"
	"case-study-checkout/repository"
	"errors"
)
type ObjProduct struct {
	Config config.Configuration
}

func NewProducts(config config.Configuration) repository.IProducts{
	return &ObjProduct{
		Config: config,
	}
}

/**
    Mockup Data
 */
func (p *ObjProduct) GetAllProducts() (arrProducts []entity.Products,err error){

	arrProducts = []entity.Products{
		{
			Sku:   "120P90",
			Name:  "Google Home",
			Price: 49.99,
			Qty:   10,
	   },
		{
			Sku:   "43N23P",
			Name:  "MacBook Pro",
			Price: 5399.99,
			Qty:   5,
		},
		{
			Sku:   "A304SD",
			Name:  "Alexa Speaker",
			Price: 109.50,
			Qty:   10,
		},
		{
			Sku:   "234234",
			Name:  "Raspberry Pi B",
			Price: 30,
			Qty:   2,
		},
	}

	return arrProducts,nil
}

func (p *ObjProduct) GetProductsBySKU(sku string) (listProducts []entity.Products,err error){

	arrProducts,_ := p.GetAllProducts()
	for _,prod := range arrProducts{
		if prod.Sku == sku{
			listProducts = append(listProducts,prod)
		}
	}
	return listProducts,errors.New("product not found")
}


