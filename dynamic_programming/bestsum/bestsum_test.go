package bestsum

import (
	"reflect"
	"testing"
)

// Test Recursion
func TestRecursionBestSum(t *testing.T) {

	x := RecursionBestSum(300, []int{7, 14, 286})

	if !reflect.DeepEqual(x, []int{286, 7, 7}) {
		t.Fatal("RecursionHowSum error")
	}
}

func BenchmarkRecursionBestSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RecursionBestSum(300, []int{7, 14})
	}
}

func TestMemoBestSum(t *testing.T) {

	x := MemoBestSum(300, []int{7, 14, 286}, make(map[int][]int))

	if !reflect.DeepEqual(x, []int{286, 7, 7}) {
		t.Fatal("RecursionBestSum error")
	}
}

func BenchmarkMemoBestSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MemoBestSum(300, []int{300, 14, 286}, make(map[int][]int))
	}
}
