package word

import (
	"fmt"
	"reflect"
	"testing"
)

func TestUseCount(t *testing.T) {
	tests := map[string]map[string]int{
		"blabla word": {
			"blabla": 1,
			"word":   1,
		},
		"blabla blabla word": {
			"blabla": 2,
			"word":   1,
		},
		"": {},
		"word": {
			"word": 1,
		},
	}

	for s, expected := range tests {
		result := UseCount(s)

		if !reflect.DeepEqual(result, expected) {
			t.Error("expected: ", expected, "got: ", result)
		}
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount("blabla blabla test text")
	}
}

func TestCount(t *testing.T) {
	tests := map[string]int{
		"blabla word":        2,
		"blabla blabla word": 2,
		"":                   0,
		"word":               1,
	}

	for s, expected := range tests {
		result := Count(s)
		if result != expected {
			t.Error("expected: ", expected, "got: ", result)
		}
	}
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count("blabla blabla test text")
	}
}

func ExampleCount() {
	fmt.Println(Count("hello there"))
	// Output: 2
}

func ExampleCount_wordDuplicate() {
	fmt.Println(Count("hello hello there"))
	// Output: 2
}
