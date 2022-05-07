package main

import "testing"

var efaceKnowType int
var efaceUnknownType uint32

/**
BenchmarkEfaceKnowType
BenchmarkEfaceKnowType/eface-to-type
BenchmarkEfaceKnowType/eface-to-type-10         	1000000000	         0.4796 ns/op
BenchmarkEfaceKnowType/int-to-int
BenchmarkEfaceKnowType/int-to-int-10            	1000000000	         0.3163 ns/op
*/
// The result above is used to test know types result.
func BenchmarkEfaceKnowType(b *testing.B) {
	b.Run("eface-to-type", func(b *testing.B) {
		var ebread interface{} = 666
		for i := 0; i < b.N; i++ {
			efaceKnowType = ebread.(int)
		}
	})

	b.Run("int-to-int", func(b *testing.B) {
		var c int32 = 666
		for i := 0; i < b.N; i++ {
			efaceKnowType = int(c)
		}
	})
}

/**
BenchmarkEfaceUnknownType
BenchmarkEfaceUnknownType/switch-small
BenchmarkEfaceUnknownType/switch-small-10         	756420516	         1.587 ns/op
BenchmarkEfaceUnknownType/switch-big
BenchmarkEfaceUnknownType/switch-big-10           	545435434	         2.197 ns/op
*/
// The result above is used to test iface unknown types result.
func BenchmarkEfaceUnknownType(b *testing.B) {
	var ebread interface{} = uint32(42)

	b.Run("switch-small", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			switch v := ebread.(type) {
			case int8:
				efaceUnknownType = uint32(v)
			case int16:
				efaceUnknownType = uint32(v)
			default:
				efaceUnknownType = v.(uint32)
			}
		}
	})
	b.Run("switch-big", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			switch v := ebread.(type) {
			case int8:
				efaceUnknownType = uint32(v)
			case int16:
				efaceUnknownType = uint32(v)
			case int32:
				efaceUnknownType = uint32(v)
			case int64:
				efaceUnknownType = uint32(v)
			case uint8:
				efaceUnknownType = uint32(v)
			case uint16:
				efaceUnknownType = uint32(v)
			case uint64:
				efaceUnknownType = uint32(v)
			default:
				efaceUnknownType = v.(uint32)
			}
		}
	})
}
