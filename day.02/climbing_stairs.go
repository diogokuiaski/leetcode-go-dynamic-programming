package day02

var memo = make(map[int]int)

func climbStairs(n int) int {
	if n <= 3 {
		return n
	}
	val, ok := memo[n]
	if ok {
		return val
	} else {
		memo[n] = climbStairs(n-1) + climbStairs(n-2)
		return memo[n]
	}
}
