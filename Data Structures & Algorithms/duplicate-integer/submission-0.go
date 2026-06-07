func hasDuplicate(nums []int) bool {
	 m := make(map[int]int)

	for _, v := range nums{
		_, ok := m[v]
		if(ok){
			// its dublicate
			return true
		}
		m[v] = v
	} 
	return false
}
