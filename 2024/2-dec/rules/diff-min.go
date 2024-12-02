package rules

var (
	Tolerance = 1
)

func MinDiff(input []int, min int) bool {
	for i := 1; i < len(input); i++ {
		if absValue(input[i]-input[i-1]) < min {
			return false
		}
	}
	return true
}

func absValue(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}
