package quicksort

func swap(data []int, i, j int) {
	temp := data[i]
	data[i] = data[j]
	data[j] = temp
}

func partition(nums []int, front, end int) ([]int, int) {
	pivot := nums[end]
	i := front - 1
	for j := front; j < end; j++ {
		if nums[j] < pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	i++
	nums[i], nums[end] = nums[end], nums[i]
	return nums, i
}

// time: avg: O(n log n), best: O(n log n), worst: O(n^2)
// space: best: O(log n), worst: O(n)
// unstable
func Quicksort(nums []int, front, end int) []int {
	if front < end {
		nums, pivot := partition(nums, front, end)
		Quicksort(nums, front, pivot-1)
		Quicksort(nums, pivot+1, end)
	}
	return nums
}
