package leetcode54

func spiralOrder(matrix [][]int) []int {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	top, left, bottom, right := 0, 0, len(matrix)-1, len(matrix[0])-1

	res := make([]int, 0)

	for {

		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
		top++
		if top > bottom {
			break
		}

		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}
		right--
		if right < left {
			break
		}

		for i := right; i >= left; i-- {
			res = append(res, matrix[bottom][i])
		}
		bottom--
		if bottom < top {
			break
		}

		for i := bottom; i >= top; i-- {
			res = append(res, matrix[i][left])
		}
		left++
		if left > right {
			break
		}
	}

	return res
}
