package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	tests := map[int]int{
		1:  7,
		2:  14,
		3:  21,
		10: 70,
	}
	for humans, dogs := range tests {
		result := Years(humans)
		if result != dogs {
			t.Error("Expected: ", dogs, " got: ", result)
		}

	}
}

func ExampleYears() {
	fmt.Println(Years(10))
	// Output:
	// 70
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(10)
	}
}

func TestYearsTwo(t *testing.T) {
	tests := map[int]int{
		1:  7,
		2:  14,
		3:  21,
		10: 70,
	}
	for humans, dogs := range tests {
		result := YearsTwo(humans)
		if result != dogs {
			t.Error("Expected: ", dogs, " got: ", result)
		}

	}
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(10))
	// Output:
	// 70
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(10)
	}
}



