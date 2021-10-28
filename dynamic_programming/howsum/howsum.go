package howsum

/*
	Write a function `HowSum(targetSum, numbers)` that takes in a targetSum and an array of numbers as arguments.

	The function should return an array containing any combination of elements that add up to exactly the targetSum.
	If there is no combination that adds up to the targetSum, then return null.

	If there are multiple combinations possible, you may return any single one.

	Examples:
	HowSum(7,[5,3,4,7]) => [3,4] or [7]
	HowSum(8,[2,3,5]) => [3,5], [2,2,2,2] or [2,3,3]
	HowSum(7,[2,4]) => null
	HowSum(0,[2,4]) => []
*/

// Recursion
// m = target sum, n = array length
// time complexity: O(n ^ m * m)
// space complexity: O(m)

func RecursionHowSum(targetSum int, numbers []int) []int {
	if targetSum == 0 {
		return []int{}
	}
	if targetSum < 0 {
		return nil
	}
	for _, num := range numbers {
		remainder := targetSum - num
		remainderResult := RecursionHowSum(remainder, numbers)
		if remainderResult != nil {
			return append(remainderResult, num)
		}
	}
	return nil
}

// Memo
// time complexity: O(m * n * m)
// space complexity: O(m * m)

func MemoHowSum(targetSum int, numbers []int, memo map[int][]int) []int {
	if val, ok := memo[targetSum]; ok {
		return val
	}
	if targetSum == 0 {
		return []int{}
	}
	if targetSum < 0 {
		return nil
	}
	for _, num := range numbers {
		remainder := targetSum - num
		remainderResult := MemoHowSum(remainder, numbers, memo)
		if remainderResult != nil {
			memo[targetSum] = append(remainderResult, num)
			return memo[targetSum]
		}
	}
	memo[targetSum] = nil
	return nil
}
