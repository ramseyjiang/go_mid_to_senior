package twotestways

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPackDoughnutsBoxSubtests(t *testing.T) {
	t.Run("It fills the box with tasty doughnuts", func(t *testing.T) {
		// Arrange
		items := []string{"Sri Lankan Cinnamon Sugar", "Mocha Tea", "Home Made Raspberry Jam", "Lime & Coconut (ve)"}
		box := newDoughnutsBox(4)

		// Act
		numOfDoughnutsInTheBox, err := box.pack(items)

		// Assert
		require.NoError(t, err)
		assert.Equal(t, numOfDoughnutsInTheBox, 4)
	})
	t.Run("It fails to fill the box with too many doughnuts", func(t *testing.T) {
		// Arrange
		items := []string{"Sri Lankan Cinnamon Sugar", "Mocha Tea", "Home Made Raspberry Jam", "Lime & Coconut (ve)", "Lime & Coconut (ve)"}
		box := newDoughnutsBox(4)

		// Act
		numOfDoughnutsInTheBox, err := box.pack(items)

		// Assert
		require.Error(t, err)
		assert.Equal(t, "failed to put 5 doughnuts in the box, it's only has 4 doughnuts capacity", err.Error())
		assert.Equal(t, 0, numOfDoughnutsInTheBox)
	})
	t.Run("It fails to put a giant chocolate cookie into the box", func(t *testing.T) {
		// Arrange
		items := []string{"Sri Lankan Cinnamon Sugar", "Giant Chocolate Cookie"}
		box := newDoughnutsBox(4)

		// Act
		numOfDoughnutsInTheBox, err := box.pack(items)

		// Assert
		require.Error(t, err)
		assert.Equal(t, "the following items cannot be placed into the box: [Giant Chocolate Cookie]", err.Error())
		assert.Equal(t, 1, numOfDoughnutsInTheBox)
	})
}
