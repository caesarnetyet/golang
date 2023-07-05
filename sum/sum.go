package sum

// Sum returns the sum of an array of numbers
func Sum(numbers []int) int {
	var sum int
	for _, v := range numbers {
		sum += v
	}

	return sum

}
