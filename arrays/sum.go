package arrays

// Returns the sum of all integers in numbers slice
func Sum(numbers []int) (sum int) {
	sum = 0
	for _, number := range numbers {
		sum += number
	}
	return
}

// Returns an integer array containing the summed values for every slice
func SumAll(numbersToSum ...[]int) (sums []int) {
	sums = make([]int, len(numbersToSum))
	for i, v := range numbersToSum {
		sums[i] = Sum(v)
	}
	return
}
