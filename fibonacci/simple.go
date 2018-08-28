package fibonacci

func Fibonacci(x int) int {
	if x == 0 {
		return 0
	} else if x <= 2 {
		return 1
	} else {
		return Fibonacci(x-2) + Fibonacci(x-1)
	}
}
