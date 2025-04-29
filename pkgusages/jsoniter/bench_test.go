package jsoniter

import (
	stdjson "encoding/json"
	"testing"
	"time"

	gojson "github.com/goccy/go-json"
	jsoniter "github.com/json-iterator/go"
)

// Declare test data structure
type User struct {
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"created_at"`
}

var (
	jsonIter = jsoniter.ConfigCompatibleWithStandardLibrary
	user     = User{
		Name:      "Alice",
		Age:       30,
		CreatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
	}
)

// Standard JSON benchmark test
func BenchmarkMarshalStdJSON(b *testing.B) {
	b.ResetTimer() // reset timer
	for i := 0; i < b.N; i++ {
		_, err := stdjson.Marshal(user)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMarshalGoJSON(b *testing.B) {
	b.ResetTimer() // reset timer
	for i := 0; i < b.N; i++ {
		_, err := gojson.Marshal(user)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// Jsoniter benchmark test
func BenchmarkMarshalJSONiter(b *testing.B) {
	b.ResetTimer() // reset timer
	for i := 0; i < b.N; i++ {
		_, err := jsonIter.Marshal(user)
		if err != nil {
			b.Fatal(err)
		}
	}
}
