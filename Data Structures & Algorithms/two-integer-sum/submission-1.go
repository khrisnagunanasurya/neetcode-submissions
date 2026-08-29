func twoSum(nums []int, target int) []int {
	memo := make(map[int]int)

	for i, value := range nums {
		diff := target - value

		if j, ok := memo[diff]; ok {
			return []int{j,i}
		}

		memo[value] = i
	}

	return nil
}
