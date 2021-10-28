package mergesort

// time: best: O(n log n), worst: O(n log n) ,avg: O(n log n)
// space: O(n)
// stable
func Mergesort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	left := Mergesort(nums[:len(nums)/2])
	right := Mergesort(nums[len(nums)/2:])
	return merge(left, right)
}
func merge(left, right []int) []int {
	final := []int{}
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			final = append(final, left[i])
			i++
		} else {
			final = append(final, right[j])
			j++
		}
	}
	if i < len(left) {
		final = append(final, left[i:]...)
	}
	if j < len(right) {
		final = append(final, right[j:]...)
	}
	return final
}
