func hasDuplicate(nums []int) bool {
	 memo := map[int]bool{}

    for _, num := range nums {
		if memo[num] {
			return true
		}

		memo[num] = true
	}

	return false
}
