func twoSum(nums []int, target int) []int {

	for i := 0; i < len(nums); i++ {
		var start = nums[i]
		for j := i + 1; j < len(nums); j++ {
			var end = nums[j]
			if end+start == target {
				return []int{i, j}
			}
		}
	}

	return []int{}

}