package day02

type costIndex struct {
	index, cost int
}

var memoMin = make(map[costIndex]int)

func minCostClimbingStairs(cost []int) int {
	memoMin = make(map[costIndex]int)
	return min(
		costClimbingStairs(cost, 0),
		costClimbingStairs(cost, 1),
	)
}

func costClimbingStairs(cost []int, index int) int {
	if index >= len(cost) {
		return 0
	}
	i := costIndex{index, cost[index]}
	val, ok := memoMin[i]
	if ok {
		return val
	} else {
		b := min(
			costClimbingStairs(cost, index+1),
			costClimbingStairs(cost, index+2),
		)
		memoMin[i] = cost[index] + b
		return memoMin[i]
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
