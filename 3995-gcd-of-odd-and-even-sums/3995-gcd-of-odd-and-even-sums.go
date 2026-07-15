func gcdOfOddEvenSums(n int) int {
	i := 1
	o := 2
	counter := 1
	for n > 1 {
		counter += 2
		i = i + counter
		o = o + counter + 1
		n--
		fmt.Println(i, o)
	}
	for o > 0 {
		fmt.Println(i, o)
		i, o = o, i%o
	}
	return i
}