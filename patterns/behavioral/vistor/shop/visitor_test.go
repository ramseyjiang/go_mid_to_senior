package shop

import (
	"testing"
)

func TestDiscounts(t *testing.T) {
	book := &Book{price: 100}
	electronic := &Electronic{price: 200}

	seasonalDiscount := &SeasonalDiscountVisitor{
		bookDiscount:       0.1,
		electronicDiscount: 0.2,
	}

	t.Run("Test Book price after discount", func(t *testing.T) {
		discountedBookPrice := book.Accept(seasonalDiscount)
		expectedBookPrice := 100 * (1 - 0.1)
		if discountedBookPrice != expectedBookPrice {
			t.Errorf("Book price after discount is incorrect, expected: %.2f, got: %.2f", expectedBookPrice, discountedBookPrice)
		}
	})

	t.Run("Test Electronic price after discount is incorrect", func(t *testing.T) {
		discountedElectronicPrice := electronic.Accept(seasonalDiscount)
		expectedElectronicPrice := 200 * (1 - 0.2)
		if discountedElectronicPrice != expectedElectronicPrice {
			t.Errorf("Electronic price after discount is incorrect, expected: %.2f, got: %.2f", expectedElectronicPrice, discountedElectronicPrice)
		}
	})
}
