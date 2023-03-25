package rotatearray

func rotate(nums []int, k int) {
	k %= len(nums)
	if len(nums) != 1 {
		slice1 := nums[:len(nums)-k]
		slice2 := nums[len(nums)-k:]
		slice2 = append(slice2, slice1...)

		copied := append(make([]int, 0, len(slice2)), slice2...)
		copy(nums, copied)
	}

}
