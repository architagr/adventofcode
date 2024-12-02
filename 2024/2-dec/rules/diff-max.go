package rules

func MaxDiff(input []int, max int) bool {
	for i := 1; i < len(input); i++ {
		if absValue(input[i]-input[i-1]) > max {
			return false
		}
	}
	return true
}
