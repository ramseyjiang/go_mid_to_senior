package transaction

import "testing"

func BenchmarkInitial(b *testing.B) {
	logs := []string{"88 99 200", "88 99 300", "99 32 100", "12 12 15"} // Can use larger dataset
	threshold := int32(2)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		processLogsInitial(logs, threshold)
	}
}

func BenchmarkOptimized(b *testing.B) {
	logs := []string{"88 99 200", "88 99 300", "99 32 100", "12 12 15"} // Can use larger dataset
	threshold := int32(2)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		processLogsOptimized(logs, threshold)
	}
}
