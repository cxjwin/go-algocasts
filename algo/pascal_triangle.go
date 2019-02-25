package algo

// https://leetcode.com/problems/pascal-triangle/

func makePascalTriangle(numRows int) [][]int {
	if numRows < 1 {
		return [][]int{}
	}

	res := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		num := i + 1
		arr := make([]int, num)
		arr[0] = 1
		if num-1 > 0 {
			arr[num-1] = 1
		}
		for j := 1; j < i; j++ {
			pre := res[i-1]
			arr[j] = pre[j-1] + pre[j]
		}
		res[i] = arr
	}

	return res
}
