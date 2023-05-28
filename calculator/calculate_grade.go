package calculator

func CalculateGrade(str string) (correct int16, wrong int16, total int16, grade float32) {
	total = int16(len(str))

	for _, char := range str {
		if char == 't' {
			correct++
		} else {
			wrong++
		}
	}

	grade = float32(correct * 100 / total)

	return

}
