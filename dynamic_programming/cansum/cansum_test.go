package cansum

import (
	"testing"
)

// Test Recursion
func TestRecursionCanSum(t *testing.T) {

	x := RecursionCanSum(7, []int{2, 4})

	if x != false {
		t.Fatal("RecursionGridTraveler error")
	}
}

func BenchmarkRecursionCanSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RecursionCanSum(300, []int{7, 14})
	}
}

func TestMemoCanSum(t *testing.T) {

	x := MemoCanSum(7, []int{2, 4}, make(map[int]bool))

	if x != false {
		t.Fatal("MemoCanSum error")
	}
}

func BenchmarkMemoCanSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MemoCanSum(300, []int{7, 14}, make(map[int]bool))
	}
}
