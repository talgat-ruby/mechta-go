package calc

func SliceMapSum(sl []map[string]int) int {
	sum := 0

	for _, m := range sl {
		for _, num := range m {
			sum += num
		}
	}

	return sum
}
