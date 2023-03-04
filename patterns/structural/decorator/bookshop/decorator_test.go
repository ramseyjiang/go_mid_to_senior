package bookshop

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestBookShop(t *testing.T) {
	book := &BookImpl{Title: "The Ramsey's Golang Pattern Guide", Price: 10.0}
	discountedBook := &DiscountedBookImpl{Book: book, Discount: 0.2}

	assert.Equal(t, book.Title, "The Ramsey's Golang Pattern Guide")
	assert.Equal(t, book.Price, book.GetPrice())

	assert.Equal(t, discountedBook.GetPrice(), 8.00)
	assert.Equal(t, discountedBook.Discount, 0.2)
}
