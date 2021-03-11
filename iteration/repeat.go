package iteration

// Returns char repeated n times
func Repeat(char string, n int) (repeated string) {
	for i := 0; i < n; i++ {
		repeated += char
	}
	return
}
