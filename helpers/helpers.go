package helpers

func IsPassed(marks int32) bool {
	passMarks := 0.3 * 1600
	return float64(marks) >= passMarks
}
