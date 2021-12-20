package day04

func jump(nums []int) int {
	visited := make(map[int]int, len(nums))
	return letsJump2(nums, visited, 0)
}

func letsJump2(nums []int, visited map[int]int, index int) int {
	if index == len(nums)-1 {
		return 0
	}

	if index > len(nums)-1 || nums[index] == 0 {
		return 99999
	}

	val, ok := visited[index]
	// fmt.Printf("testing %v | %v index %d (%d) \n", nums, visited, nums[index], index)
	if ok {
		// fmt.Printf("val :: %d (%d)\n", val, index)
		return val
	} else {
		acc := 99999
		for i := index + 1; i <= index+nums[index]; i++ {
			r := letsJump2(nums, visited, i)
			acc = min(acc, r)
			// fmt.Printf("add :: %d (%d) :: %d/%d\n", nums[i], i, acc, r)
		}
		visited[index] = acc + 1
		return visited[index]
	}
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func jump2Best(nums []int) int {
	n, steps, lastPerimeter, perimeter := len(nums)-1, 0, 0, 0

	for i := 0; i < n; i++ {
		contender := i + nums[i]
		if perimeter < contender {
			if contender >= n { // can reach the last index by one jump
				return steps + 1
			}
			perimeter = contender
		}
		if i == lastPerimeter {
			steps++
			lastPerimeter = perimeter
		}
	}
	return steps
}
