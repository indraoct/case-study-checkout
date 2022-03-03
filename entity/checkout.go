package entity


type CheckoutTotal struct {
	Checkout Checkout `json:"checkout"`
	Promos  []Promo   `json:"promos"`
}