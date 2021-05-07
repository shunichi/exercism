// Package diffsquares calculates the difference between the square of the sum and the sum of the squares of the first N natural numbers
package diffsquares

// SquareOfSum calculates the square of the sum of the first N natural numbers
func SquareOfSum(n int) int {
	tmp := n * (n + 1) / 2
	return tmp * tmp
}

// SumOfSquares calculates the sum of the square of the first N natural numbers
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

// Difference calculates the difference between the square of the sum and the sum of the squares of the first N natural numbers
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
