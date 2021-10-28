package selectionsort

// time: best: O(n^2), worst: O(n^2), avg: O(n^2)
// space: O(1)
// stable
func Selectionsort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		min := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		if min != i {
			nums[i], nums[min] = nums[min], nums[i]
		}
	}
	return nums
}
