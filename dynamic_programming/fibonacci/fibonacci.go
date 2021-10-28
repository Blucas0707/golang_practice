package fabonacci

// Recursion
// time complexity: O(2^N)
// space complexity: O(N)

func RecursionFib(n int) int {
	if n <= 2 {
		return 1
	}
	return RecursionFib(n-1) + RecursionFib(n-2)
}

// Memorization
// time complexity: O(N)
// space complexity: O(N)

func MemoFib(n int, memo map[int]int) int {
	if val, ok := memo[n]; ok {
		return val
	}
	if n <= 2 {
		return 1
	}
	memo[n] = MemoFib(n-1, memo) + MemoFib(n-2, memo)
	return memo[n]
}
