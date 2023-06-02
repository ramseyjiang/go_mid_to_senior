package sort

import (
	"math/rand"
	"testing"
	"time"
)

const arrayNum = 10000
const arrayMaxNum = 100000

func generateUnsortedArr() []int {
	rand.Seed(time.Now().Unix()) // It is used to confirm rand() will generate a random number each run time.
	arr := make([]int, arrayNum)
	for i := 0; i <= arrayNum-1; i++ {
		arr[i] = rand.Intn(arrayMaxNum)
	}
	return arr
}

func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := generateUnsortedArr()
		quickSort(arr)
	}
}

func BenchmarkMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := generateUnsortedArr()
		mergeSort(arr)
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := generateUnsortedArr()
		bubbleSort(arr)
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := generateUnsortedArr()
		insertionSort(arr)
	}
}

func BenchmarkSelectionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := generateUnsortedArr()
		selectionSort(arr)
	}
}
