package gridtraveler

import "strconv"

/*
	Say that you are a traveler on a 2D grid.
	You begin in the top-left corner and your goal is to travel to the bottom-right corner.
	You may only move down or right.

	In how many ways can you travel to the goal on a grid with dimensions m * n?
	Write a function `GridTraveler(m, n)` that cakculates this.
	example:
	GridTraveler(2, 3) => 3 ways
	[[0,0,0],
	 [0,0,0]]
*/

// Recursion
// time complexity: O(2^(m + n))
// space complexity: O(m + n)

func RecursionGridTraveler(m, n int) int {
	if m == 0 || n == 0 {
		return 0
	}
	if m == 1 && n == 1 {
		return 1
	}
	return RecursionGridTraveler(m-1, n) + RecursionGridTraveler(m, n-1)
}

// Memo
// time complexity: O(m * n)
// space complexity: O(m + n)

func MemoGridTraveler(m int, n int, memo map[string]int) int {
	key := strconv.Itoa(m) + "," + strconv.Itoa(n)
	if val, ok := memo[key]; ok {
		return val
	}
	if m == 0 || n == 0 {
		return 0
	}
	if m == 1 && n == 1 {
		return 1
	}
	memo[key] = MemoGridTraveler(m-1, n, memo) + MemoGridTraveler(m, n-1, memo)
	return memo[key]
}
