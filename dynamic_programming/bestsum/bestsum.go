package bestsum

/*
	Write a function `BestSum(targetSum, numbers)` that takes in a targetSum and an array of numbers as arguments.

	The function should return an array containing the shortest combination of numbers that add up to exactly the targetSum.
	If there is no combination that adds up to the targetSum, then return null.

	If there is a tie for the shortest combination, you may return any one of the shortest.

	Examples:
	BestSum(7,[5,3,4,7]) => [3,4] or [7]
	BestSum(8,[2,3,5]) => [3,5], [2,2,2,2] or [2,3,3]
	BestSum(7,[2,4]) => null
	BestSum(0,[2,4]) => []
*/

// Recursion
// m = target sum, n = array length
// time complexity: O(n ^ m * m)
// space complexity: O(m)

func RecursionBestSum(targetSum int, numbers []int) []int {
	if targetSum == 0 {
		return []int{}
	}
	if targetSum < 0 {
		return nil
	}
	for _, num := range numbers {
		remainder := targetSum - num
		remainderResult := RecursionBestSum(remainder, numbers)
		if remainderResult != nil {
			return append(remainderResult, num)
		}
	}
	return nil
}

// Memo
// time complexity: O(m * n * m)
// space complexity: O(m * m)

func MemoBestSum(targetSum int, numbers []int, memo map[int][]int) []int {
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
		remainderResult := MemoBestSum(remainder, numbers, memo)
		if remainderResult != nil {
			memo[targetSum] = append(remainderResult, num)
			return memo[targetSum]
		}
	}
	memo[targetSum] = nil
	return nil
}
