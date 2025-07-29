package main

// Change these boolean values to control whether you see
// the expected answer and/or hints.
const showExpectedResult = false
const showHints = false

type CartItem struct {
	Name     string
	Price    float64
	Quantity int
}

// calculateTotal() returns the total value of the shopping cart.
func calculateTotal(cart []CartItem) float64 {
	var res float64 = 0
	for i := 0; i < len(cart); i++ {
		res += cart[i].Price * float64(cart[i].Quantity)
	}
	return res
}
