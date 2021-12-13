package day03

func rob(nums []int) int {
	memoRob := make(map[RobIndex]int)
	return Max(
		maxRob(nums, 0, memoRob),
		maxRob(nums, 1, memoRob),
	)
}

func maxRob(cost []int, index int, memoRob map[RobIndex]int) int {
	if index >= len(cost) {
		return 0
	}
	i := RobIndex{index, cost[index]}
	val, ok := memoRob[i]
	if ok {
		return val
	} else {
		b := Max(
			maxRob(cost, index+2, memoRob),
			maxRob(cost, index+3, memoRob),
		)
		memoRob[i] = cost[index] + b
		return memoRob[i]
	}
}
