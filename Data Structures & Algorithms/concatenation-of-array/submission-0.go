func getConcatenation(nums []int) []int {
    return append(append([]int{}, nums...), nums...)
}
