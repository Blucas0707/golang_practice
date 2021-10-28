package cansum

/*
	Write a function `CanSum(targetSum, numbers)` that takes in a targetSum and an array of numbers as arguments.

	The function should return a boolean indicatiing whether or not it is possible to generate the targetSum using numbers from the array.

	You may use an element of the array as many times as needed.

	You may assume that all input numbers are nonnegative

	Examples:
	CanSum(7,[5,3,4,7]) => true
	CanSum(7,[2,4]) => false
*/

// Recursion
// m = target sum, n = array length
// time complexity: O(n ^ m)
// space complexity: O(m)

func RecursionCanSum(targetSum int, numbers []int) bool {
	if targetSum == 0 {
		return true
	}
	if targetSum < 0 {
		return false
	}
	for _, number := range numbers {
		remainder := targetSum - number
		if RecursionCanSum(remainder, numbers) {
			return true
		}
	}
	return false
}

// Memo
// time complexity: O(m * n)
// space complexity: O(m)

func MemoCanSum(targetSum int, numbers []int, memo map[int]bool) bool {
	if val, ok := memo[targetSum]; ok {
		return val
	}
	if targetSum == 0 {
		memo[targetSum] = true
		return memo[targetSum]
	}
	if targetSum < 0 {
		memo[targetSum] = false
		return memo[targetSum]
	}
	for _, number := range numbers {
		remainder := targetSum - number
		if MemoCanSum(remainder, numbers, memo) {
			memo[targetSum] = true
			return true
		}
	}
	memo[targetSum] = false
	return false
}
