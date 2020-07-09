package diffsquares

func SquareOfSum(n int) int {
	sum := (1 + n) * n / 2
	return sum * sum
}

func SumOfSquares(n int) int {
	result := 0
	for i := 1; i <= n; i++ {
		result += i * i
	}
	return result
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
