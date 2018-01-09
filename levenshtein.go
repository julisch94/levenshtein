package levenshtein

func calculateDistance(firstWord string, secondWord string) int {
	// Create matrix
	numOfColumns := len(firstWord) + 1
	numOfRows := len(secondWord) + 1
	distanceMatrix := createZeroedMatrix(numOfColumns, numOfRows)

	// Prepare matrix
	distanceMatrix = prepareMatrix(distanceMatrix)

	// Run algorithmus
	distanceMatrix = runLevenshteinAlgorithmOnMatrix(distanceMatrix, firstWord, secondWord)

	// result distance is in the bottom right cell
	levenshteinDistance := distanceMatrix[numOfColumns-1][numOfRows-1]
	return levenshteinDistance
}

func createZeroedMatrix(numOfColumns int, numOfRows int) [][]int {
	matrix := make([][]int, numOfColumns)

	// Fill every column with empty arrays
	for columnIndex := 0; columnIndex < numOfColumns; columnIndex++ {
		matrix[columnIndex] = make([]int, numOfRows)
	}

	return matrix
}

func prepareMatrix(distanceMatrix [][]int) [][]int {
	numOfColumns := getNumberOfMatrixColumns(distanceMatrix)
	numOfRows := getNumberOfMatrixRows(distanceMatrix)

	// Fill first row with 0,1,2,...
	for columnIndex := 0; columnIndex < numOfColumns; columnIndex++ {
		distanceMatrix[columnIndex][0] = columnIndex
	}
	// Fill first column with 0,1,2,...
	for rowIndex := 0; rowIndex < numOfRows; rowIndex++ {
		distanceMatrix[0][rowIndex] = rowIndex
	}

	return distanceMatrix
}

func runLevenshteinAlgorithmOnMatrix(distanceMatrix [][]int, firstWord string, secondWord string) [][] int {
	numOfColumns := getNumberOfMatrixColumns(distanceMatrix)
	numOfRows := getNumberOfMatrixRows(distanceMatrix)

	for rowIndex := 1; rowIndex < numOfRows; rowIndex++ {
		for columnIndex := 1; columnIndex < numOfColumns; columnIndex++ {
			if firstWord[columnIndex-1] == secondWord[rowIndex-1] {
				distanceMatrix[columnIndex][rowIndex] = distanceMatrix[columnIndex-1][rowIndex-1]
			} else {
				costOfSubstitution := distanceMatrix[columnIndex-1][rowIndex-1] + 1
				costOfInsertion := distanceMatrix[columnIndex][rowIndex-1] + 1
				costOfDeletion := distanceMatrix[columnIndex-1][rowIndex] + 1

				costOfCheapest := getMinimumOfNumbers(costOfSubstitution, costOfInsertion, costOfDeletion)
				distanceMatrix[columnIndex][rowIndex] = costOfCheapest
			}
		}
	}

	return distanceMatrix
}
