func arrayRankTransform(arr []int) []int {
	numToRank := make(map[int]int)

	sortedArr := make([]int, len(arr))

	copy(sortedArr, arr)
	sort.Ints(sortedArr)
	rank := 1

	for i, v := range sortedArr {
		if i > 0 && sortedArr[i] > sortedArr[i-1] {
			rank++
		}
		numToRank[v] = rank
	}
result := make([]int, len(arr))
	for i, v := range arr {
		result[i] = numToRank[v]
	}

	return result 
    }