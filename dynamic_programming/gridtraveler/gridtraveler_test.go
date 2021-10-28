package gridtraveler

import "testing"

// Test Recursion
func TestRecursionGridTraveler(t *testing.T) {

	x := RecursionGridTraveler(10, 10)

	if x != 48620 {
		t.Fatal("RecursionGridTraveler error")
	}
}

func BenchmarkRecursionGridTraveler(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RecursionGridTraveler(10, 10)
	}
}

func TestMemoGridTraveler(t *testing.T) {

	x := MemoGridTraveler(10, 10, make(map[string]int))

	if x != 48620 {
		t.Fatal("MemoFib error")
	}
}

func BenchmarkMemoGridTraveler(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MemoGridTraveler(10, 10, make(map[string]int))
	}
}
