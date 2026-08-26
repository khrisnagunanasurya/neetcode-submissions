func hasDuplicate(nums []int) bool {
	 memo := make(map[int]struct{}, len(nums))

    for _, num := range nums {
		if _, exists := memo[num]; exists {
			return true
		}

		memo[num] = struct{}{}
	}

	return false
}
