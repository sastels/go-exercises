package diffsquares

// SquareOfSum squares the sum of the first n integers
func SquareOfSum(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum * sum
}

// SumOfSquares sums the squares of the first n integers
func SumOfSquares(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i * i
	}
	return sum
}

// Difference subtracts the two
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
