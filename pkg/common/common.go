package common

func ColNotEmpty(row []string) int {
	count := 0
	for _, col := range row {
		if col != "" {
			count++
		}
	}
	return count
}
