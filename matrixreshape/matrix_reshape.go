package matrixreshape

func matrixReshape(mat [][]int, r int, c int) [][]int {
	rows := len(mat)
	cols := len(mat[0])
	if rows*cols != r*c {
		return mat
	}
	slice := make([]int, 0)
	for _, v := range mat {
		for _, i := range v {
			slice = append(slice, i)
		}
	}
	k := 0
	result := make([][]int, r)
	for i := 0; i < r; i++ {
		result[i] = make([]int, c)
		for j := 0; j < c; j++ {
			result[i][j] = slice[k]
			k++
		}
	}

	return result
}
