package ecommerce

import "testing"

func TestCart(t *testing.T) {
	// Given
	t.Run("Given a cart with added products", func(t *testing.T) {
		cart := &Cart{}

		// When
		t.Run("When the user adds a product to the cart", func(t *testing.T) {
			product := Product{Name: "Go Book", Price: 10.0}
			cart.AddProduct(product)

			// Then
			if len(cart.Products) != 1 {
				t.Errorf("Expected 1 product, but got %d", len(cart.Products))
			}

			if cart.Products[0] != product {
				t.Errorf("Expected %v, but got %v", product, cart.Products[0])
			}

			// And
			t.Run("And the total price should be updated accordingly", func(t *testing.T) {
				if cart.TotalPrice() != product.Price {
					t.Errorf("Expected total price to be %f, but got %f", product.Price, cart.TotalPrice())
				}
			})
		})
	})
}
