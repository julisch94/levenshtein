package levenshtein

func getNumberOfMatrixColumns(matrix [][]int) int {
	return len(matrix)
}

func getNumberOfMatrixRows(matrix [][]int) int {
	return len(matrix[0])
}

func getMinimumOfNumbers(numbers ...int) int {
	if len(numbers) == 0 {
		panic("no parameters supplied")
	}

	currentMinimum := numbers[0]
	for _, number := range numbers {
		if number < currentMinimum {
			currentMinimum = number
		}
	}

	return currentMinimum
}
