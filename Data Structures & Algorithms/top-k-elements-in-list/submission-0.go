func topKFrequent(nums []int, k int) []int {
	// 1. считаем частоты
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}

	// 2. переносим в слайс
	type pair struct {
		num  int
		freq int
	}

	var arr []pair
	for n, f := range freq {
		arr = append(arr, pair{n, f})
	}

	// 3. сортируем по частоте
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].freq > arr[j].freq
	})

	// 4. берём top k
	result := make([]int, k)
	for i := 0; i < k; i++ {
		result[i] = arr[i].num
	}

	return result
}
