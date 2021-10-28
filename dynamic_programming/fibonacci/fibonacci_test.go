package fabonacci

import "testing"

// Test Recursion
func TestRecursionFib(t *testing.T) {

	x := RecursionFib(30)

	if x != 832040 {
		t.Fatal("RecursionFib error")
	}
}

func BenchmarkRecursionFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RecursionFib(30)
	}
}

func TestMemoFib(t *testing.T) {

	x := MemoFib(30, make(map[int]int))

	if x != 832040 {
		t.Fatal("MemoFib error")
	}
}

func BenchmarkMemoFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MemoFib(30, make(map[int]int))
	}
}
