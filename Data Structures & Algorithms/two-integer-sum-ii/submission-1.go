func twoSum(numbers []int, target int) []int {
	f, l := 0, len(numbers) - 1
	for ; f < l ;{
		fe := numbers[f]
		le := numbers[l]
		if fe + le > target{
			l--
			continue
		}else if fe + le == target {
			return []int{f + 1, l + 1}
		}
		f++
	}
	return nil
}
