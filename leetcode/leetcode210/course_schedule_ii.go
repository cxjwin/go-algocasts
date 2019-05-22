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
	temp := make([]int, 0)
	res := &temp

	for i := 0; i < numCourses; i++ {
		visited := make([]bool, numCourses)
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

func findOrderTopoSort(numCourses int, prerequisites [][]int) []int {
	if numCourses <= 1 || prerequisites == nil || len(prerequisites) == 0 {
		return []int{}
	}

	graph := make([][]int, numCourses)
	for i := 0; i < numCourses; i++ {
		graph[i] = make([]int, 0)
	}

	inDegrees := make([]int, numCourses)
	for _, v := range prerequisites {
		arr := graph[v[1]]
		arr = append(arr, v[0])
		graph[v[1]] = arr

		inDegrees[v[0]]++
	}

	st := make([]int, 0)
	for _, v := range inDegrees {
		if v == 0 {
			st = append(st, v)
		}
	}

	output := make([]int, 0)

	for len(st) != 0 {
		idx := st[len(st)-1]
		st = st[:len(st)-1]

		output = append(output, idx)

		arr := graph[idx]
		for _, v := range arr {
			inDegrees[v]--
			if inDegrees[v] == 0 {
				st = append(st, v)
			}
		}
	}

	return output
}
