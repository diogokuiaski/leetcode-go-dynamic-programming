package day01

var cache = make(map[int]int)

func tribonacci(n int) int {
	if n == 2 {
		return 1
	}
	if n <= 1 {
		return n
	}

	value, ok := cache[n]
	if ok {
		return value
	}

	cache[n] = tribonacci(n-1) + tribonacci(n-2) + tribonacci(n-3)
	return cache[n]
}
