package mymath

import (
	"fmt"
	"testing"
)

func TestCenteredAvg(t *testing.T) {
	tests := map[float64][]int{
		3:   {1, 2, 4, 9},
		0.5: {0, 1, 4, 0},
		2.5: {-9, 7, 5, 0},
	}

	for expected, input := range tests {
		avg := CenteredAvg(input)
		if avg != expected {
			t.Error("Expected: ", expected, " got: ", avg)
		}
	}
}

func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{1, 2, 4, 9}))
	// Output: 3
}

func ExampleCenteredAvg_withNegatives() {
	fmt.Println(CenteredAvg([]int{-9, 7, 5, 0}))
	// Output: 2.5
}

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{1, 2, 4, 9})
	}
}
