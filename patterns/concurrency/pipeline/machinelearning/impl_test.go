package machinelearning

import (
	"reflect"
	"testing"
)

func TestPipelineIntegration(t *testing.T) {
	inputData := "the quick brown fox jumps over the lazy dog and keeps running"
	in := make(chan string, 1)
	in <- inputData
	close(in)

	vector := <-vectorised(stem(removeStopWords(tokenize(in))))

	expectedVector := []float64{0.1, 0.2, 0.3} // As defined in our vectorised function
	if !reflect.DeepEqual(vector, expectedVector) {
		t.Errorf("got %v, want %v", vector, expectedVector)
	}
}

func TestPipelineStages(t *testing.T) {
	t.Run("Tokenization", func(t *testing.T) {
		in := make(chan string, 1)
		in <- "the quick brown fox"
		close(in)

		got := <-tokenize(in)
		want := []string{"the", "quick", "brown", "fox"}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Stop-word Removal", func(t *testing.T) {
		in := make(chan []string, 1)
		in <- []string{"the", "quick", "brown", "fox"}
		close(in)

		got := <-removeStopWords(in)
		want := []string{"quick", "brown", "fox"}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Stemming", func(t *testing.T) {
		in := make(chan []string, 1)
		in <- []string{"running", "flies"}
		close(in)

		got := <-stem(in)
		want := []string{"run", "fly"}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Vectorization", func(t *testing.T) {
		in := make(chan []string, 1)
		in <- []string{"dummy", "data"} // Input doesn't affect vectorize's current behavior
		close(in)

		got := <-vectorised(in)
		want := []float64{0.1, 0.2, 0.3} // This is the static output we expect

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
