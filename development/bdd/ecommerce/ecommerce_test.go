package ecommerce

import "testing"

func TestCart(t *testing.T) {
	// Given
	t.Run("Given a user with items in their shopping cart", func(t *testing.T) {
		cart := &Cart{}
		product := Product{Name: "Go Book", Price: 10.0}
		cart.AddProduct(product)

		// When
		t.Run("When the user goes to the checkout page and inputs valid payment and shipping information", func(t *testing.T) {
			cart.Checkout()

			// Then
			t.Run("Then their purchase is confirmed and the items are removed from their shopping cart", func(t *testing.T) {
				if cart.TotalItems() != 0 {
					t.Errorf("Expected no items in cart, but got %d", cart.TotalItems())
				}
			})
		})
	})
}
