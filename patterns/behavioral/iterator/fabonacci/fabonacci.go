package fabonacci

// Iterator interface defines the methods for iterating over the sequence:
// HasNext() to check if there are more elements in the sequence,
// and Next() to get the next number in the sequence.
type Iterator interface {
	HasNext() bool
	Next() int
}

type Aggregate interface {
	Iterator() Iterator
}

// Fibonacci struct holds the first two numbers in the Fibonacci sequence (a and b) and the maximum number of iterations (max).
type Fibonacci struct {
	a, b, max int
}

// FibonacciIterator implements this interface and maintains the state of the iteration,
// keeping track of the current number in the sequence and updating the values of a and b as necessary.
type FibonacciIterator struct {
	f   *Fibonacci
	pos int
}

// New method returns a new Iterator for the sequence.
func (f *Fibonacci) New() Iterator {
	return &FibonacciIterator{f: f, pos: 0}
}

// HasNext compare pos with max, if pos less than max, it will return true. It means it has the next one.
func (fi *FibonacciIterator) HasNext() bool {
	return fi.pos < fi.f.max
}

func (fi *FibonacciIterator) Next() int {
	if fi.pos == 0 {
		fi.pos++
		return fi.f.a
	}

	if fi.pos == 1 {
		fi.pos++
		return fi.f.b
	}

	result := fi.f.a + fi.f.b
	fi.f.a = fi.f.b
	fi.f.b = result
	fi.pos++
	return result
}
