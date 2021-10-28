package howsum

import (
	"reflect"
	"testing"
)

// Test Recursion
func TestRecursionHowSum(t *testing.T) {

	x := RecursionHowSum(300, []int{7, 14, 286})

	if !reflect.DeepEqual(x, []int{286, 7, 7}) {
		t.Fatal("RecursionHowSum error")
	}
}

func BenchmarkRecursionHowSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RecursionHowSum(300, []int{7, 14})
	}
}

func TestMemoHowSum(t *testing.T) {

	x := MemoHowSum(300, []int{7, 14, 286}, make(map[int][]int))

	if !reflect.DeepEqual(x, []int{286, 7, 7}) {
		t.Fatal("RecursionHowSum error")
	}
}

func BenchmarkMemoHowSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MemoHowSum(300, []int{300, 14, 286}, make(map[int][]int))
	}
}
