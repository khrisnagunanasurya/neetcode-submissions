func twoSum(nums []int, target int) []int {
	memo := make(map[int]int, len(nums))

	for i, value := range nums {
		if j, ok := memo[target - value]; ok {
			return []int{j,i}
		}

		memo[value] = i
	}

	return nil
}
