package day01

var memo = make(map[int]int)

func fib(n int) int {
	if n <= 1 {
		return n
	}
	val, ok := memo[n]
	if ok {
		return val
	} else {
		memo[n] = fib(n-1) + fib(n-2)
		return memo[n]
	}
}
