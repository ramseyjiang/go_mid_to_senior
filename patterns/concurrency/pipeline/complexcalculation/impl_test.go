package complexcalculation

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestComplexCalculation(t *testing.T) {
	t.Run("TestGenerator", func(t *testing.T) {
		out := generator(5)
		for i := 1; i <= 5; i++ {
			if got := <-out; got != i {
				t.Errorf("generator(5) = %d; want %d", got, i)
			}
		}
	})

	t.Run("TestPower", func(t *testing.T) {
		in := make(chan int, 5)
		for i := 1; i <= 5; i++ {
			in <- i
		}
		close(in)

		out := power(in)
		for i := 1; i <= 5; i++ {
			want := i * i
			if got := <-out; got != want {
				t.Errorf("power(%d) = %d; want %d", i, got, want)
			}
		}
	})

	t.Run("TestSum", func(t *testing.T) {
		in := make(chan int, 5)
		in <- 1
		in <- 2
		in <- 3
		close(in)

		out := sum(in)
		if got := <-out; got != 6 {
			t.Errorf("sum = %d; want 6", got)
		}
	})

	t.Run("TestLaunchPipeline", func(t *testing.T) {
		tableTest := [][]int{
			{3, 14},
			{5, 55},
		}

		var res int
		for _, test := range tableTest {
			res = LaunchPipeline(test[0])
			if res != test[1] {
				t.Fatal()
			}
			assert.Equal(t, res, test[1])
		}
	})
}
