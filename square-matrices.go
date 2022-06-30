package l33t

type Matrix [][]int

// IsXMatrix returns true if the matrix is a X matrix.
func IsXMatrix(matrix Matrix) bool {
	if len(matrix) != len(matrix[0]) {
		return false
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] != 1 {
				return false
			}
		}
	}

	return true
}
