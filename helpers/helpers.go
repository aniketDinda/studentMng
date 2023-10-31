package helpers

func IsPassed(marks int32) bool {
	passMarks := 0.3 * 1600
	return float64(marks) >= passMarks
}

func GetRank(marksArray []int32, studentMark int32) int {
	rank := 1
	for _, marks := range marksArray {
		if marks > studentMark {
			rank++
		} else {
			break
		}
	}
	return rank
}
