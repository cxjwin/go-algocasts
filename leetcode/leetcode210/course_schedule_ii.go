package leetcode210

// https://leetcode.com/problems/course-schedule-ii/

func hasCycle(graph map[int][]int, checked []bool, visited []bool, order *[]int, v int) bool {
	if visited[v] {
		return true
	}

	visited[v] = true

	for _, i := range graph[v] {
		if !checked[i] && hasCycle(graph, checked, visited, order, i) {
			return true
		}
	}

	*order = append(*order, v)
	checked[v] = true
	visited[v] = false

	return false
}

func findOrder(numCourses int, prerequisites [][]int) []int {
	if numCourses <= 1 || prerequisites == nil || len(prerequisites) == 0 {
		return []int{}
	}

	m := make(map[int][]int)
	for _, v := range prerequisites {
		arr, ok := m[v[1]]
		if ok {
			arr = append(arr, v[0])
		} else {
			arr = []int{v[0]}
		}
		m[v[1]] = arr
	}

	checked := make([]bool, numCourses)
	visited := make([]bool, numCourses)
	temp := make([]int, 0)
	res := &temp

	for i := 0; i < numCourses; i++ {
		if !checked[i] && hasCycle(m, checked, visited, res, i) {
			return []int{}
		}
	}

	n := len(*res)
	output := make([]int, n)
	for i := 0; i < n; i++ {
		output[n-1-i] = (*res)[i]
	}

	return output
}
