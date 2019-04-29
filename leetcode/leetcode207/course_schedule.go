package leetcode207

// https://leetcode.com/problems/course-schedule/

func hasCycle(graph map[int][]int, checked []bool, visited []bool, v int) bool {
	if visited[v] {
		return true
	}

	visited[v] = true

	for _, i := range graph[v] {
		if !checked[i] && hasCycle(graph, checked, visited, i) {
			return true
		}
	}

	checked[v] = true
	visited[v] = false

	return false
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	if numCourses <= 1 || prerequisites == nil || len(prerequisites) == 0 {
		return true
	}

	m := make(map[int][]int)

	for _, pair := range prerequisites {
		arr, ok := m[pair[1]]
		if !ok {
			arr = []int{pair[0]}
		} else {
			arr = append(arr, pair[0])
		}
		m[pair[1]] = arr
	}

	checked := make([]bool, numCourses)
	visited := make([]bool, numCourses)

	for i := 0; i < numCourses; i++ {
		if !checked[i] && hasCycle(m, checked, visited, i) {
			return false
		}
	}

	return true
}

func canFinishTopoSort(numCourses int, prerequisites [][]int) bool {
	if numCourses <= 1 || prerequisites == nil || len(prerequisites) == 0 {
		return true
	}

	inDegree := make([]int, numCourses)

	graph := make([][]int, numCourses)
	for i := 0; i < numCourses; i++ {
		graph[i] = make([]int, 0)
	}

	for _, v := range prerequisites {
		arr := graph[v[1]]
		arr = append(arr, v[0])
		graph[v[1]] = arr
		inDegree[v[0]]++
	}

	queue := make([]int, 0)
	for i, v := range inDegree {
		if v == 0 {
			queue = append(queue, i)
		}
	}

	count := 0

	for len(queue) != 0 {
		idx := queue[len(queue)-1]
		queue = queue[:len(queue)-1]

		count++
		for _, i := range graph[idx] {
			inDegree[i]--
			if inDegree[i] == 0 {
				queue = append(queue, i)
			}
		}
	}

	return count == numCourses
}
