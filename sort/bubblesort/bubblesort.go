package bubblesort

// time: best: O(n), worst: O(n^2), avg: O(n^2)
// space: O(1)
// stable
func Bubblesort(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j+1] < nums[j] {
				nums[j+1], nums[j] = nums[j], nums[j+1]
			}
		}
	}
	return nums
}
