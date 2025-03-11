package commonwordnobanned

import "testing"

func TestMostCommonWord(t *testing.T) {
	tests := []struct {
		name      string
		paragraph string
		banned    []string
		want      string
	}{
		{
			name:      "normal case",
			paragraph: "Bob hit a ball, the hit BALL flew far after it was hit.",
			banned:    []string{"hit"},
			want:      "ball",
		},
		{
			name:      "empty paragraph",
			paragraph: "",
			banned:    []string{"test"},
			want:      "",
		},
		{
			name:      "all words banned",
			paragraph: "a a a b b c",
			banned:    []string{"a", "b", "c"},
			want:      "",
		},
		{
			name:      "case insensitive check",
			paragraph: "Apple apple APPLE",
			banned:    []string{"banana"},
			want:      "apple",
		},
		{
			name:      "punctuation handling",
			paragraph: "Hello! world... hello? ,world;",
			banned:    []string{},
			want:      "hello",
		},
		{
			name:      "tie breaker",
			paragraph: "cat dog cat dog bird",
			banned:    []string{"bird"},
			want:      "cat", // 任意选择一个即可，具体取决于实现
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mostCommonWord(tt.paragraph, tt.banned); got != tt.want {
				t.Errorf("mostCommonWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
