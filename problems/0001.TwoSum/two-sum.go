package problems

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))

	for i, num := range nums {
		if j, ok := m[target - num]; ok {
			return []int{j, i}
		}

		// 把值和序号存入map中
		m[num] = i
	}

	return nil
}