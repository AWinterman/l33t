package l33t

func RunningSum(nums []int) (sums []int) {
	sums = make([]int, len(nums))
	for i, n := range nums {
		previous := 0
		if i > 0 {
			previous = sums[i-1]
		}
		sums[i] = n + previous
	}
	return sums

}
