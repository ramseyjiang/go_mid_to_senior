package main

import (
	"log"
	"math"
	"sort"
)

func main() {
	sortIntSlice()
	sortFloatSlice()
	sortCustomiseSlice()
	sortStringsByLen()
	sortMaps()
}

// sort.IntsAreSorted() is used to check whether an int slice is sorted or not.
func sortIntSlice() {
	intSlice := []int{4, 5, 2, 1, 3, 9, 7, 8, 6, 10}
	log.Println(sort.IntsAreSorted(intSlice)) // false
	sort.Ints(intSlice)
	log.Println(intSlice)                     // [1 2 3 4 5 6 7 8 9 10]
	log.Println(sort.IntsAreSorted(intSlice)) // true
}

// sort.Float64s is used to check whether a float64 slice is sorted or not.
func sortFloatSlice() {
	// math.NaN() returns an IEEE 754 “not-a-number” value.
	floatArray := []float64{math.NaN(), -0.25, -1.32, 0.92, 4.812, 2.111}

	// Float64sAreSorted reports whether a slice is sorted in increasing order, with not-a-number (NaN) values before any other values.
	log.Println(sort.Float64sAreSorted(floatArray)) // false

	sort.Float64s(floatArray)
	log.Println(floatArray)                         // [NaN -1.32 -0.25 0.92 2.111 4.812]
	log.Println(sort.Float64sAreSorted(floatArray)) // true
}

type Person struct {
	name string
	age  int
}

var (
	personList = []Person{
		{"Olivia", 20},
		{"Jeremy", 40},
		{"Thomas", 30},
		{"Judith", 31},
		{"Angie", 33},
	}
)

// sortCustomiseSlice shows Sort a customised slice by an element
func sortCustomiseSlice() {
	sort.Slice(personList, func(i, j int) bool {
		return personList[i].age < personList[j].age
	})

	// Sorted by age:  [{Olivia 20} {Thomas 30} {Judith 31} {Angie 33} {Jeremy 40}]
	log.Println("Sorted by age: ", personList)

	// Sorting by Person by Name
	sort.Slice(personList, func(i, j int) bool {
		return personList[i].name < personList[j].name
	})

	// Sorted by name:  [{Angie 33} {Jeremy 40} {Judith 31} {Olivia 20} {Thomas 30}]
	log.Println("Sorted by name: ", personList)
}

type LengthBasedStrings []string

func (s LengthBasedStrings) Len() int {
	return len(s)
}
func (s LengthBasedStrings) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s LengthBasedStrings) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

// sortStringsByLen is used sort strings by length.
func sortStringsByLen() {
	words := []string{"rust-lang", "golang", "java", "c", "c++", "dot-net", "perl"}
	sort.Sort(LengthBasedStrings(words))

	// Sorted by Length: [c c++ perl java golang dot-net rust-lang]
	log.Println("Sorting by Length:", words)
}

// sortMaps can be used sort map by key or value.
func sortMaps() {
	items := map[string]int{
		"coin":   12,
		"chair":  3,
		"pen":    4,
		"bottle": 9,
	}

	newItems := make(map[int]string)
	keys := make([]string, 0, len(items))
	values := make([]int, 0, len(items))
	for key, value := range items {
		keys = append(keys, key)
		values = append(values, value)
		newItems[value] = key
	}

	sort.Strings(keys)
	// log.Println("Sorting by Length:", keys) // Here will be output keys only.
	for _, key := range keys {
		log.Printf("%s %d\n", key, items[key])
	}

	sort.Ints(values)
	log.Println("Sorting by Length:", values) // Here will be output keys only.
	for _, value := range values {
		log.Printf("%d %s\n", value, newItems[value])
	}
}
