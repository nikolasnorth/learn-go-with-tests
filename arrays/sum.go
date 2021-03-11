package arrays

// Returns the sum of all integers in numbers slice
func Sum(numbers []int) (sum int) {
	sum = 0
	for _, number := range numbers {
		sum += number
	}
	return
}
