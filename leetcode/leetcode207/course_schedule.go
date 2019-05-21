package leetcode207

// https://leetcode.com/problems/course-schedule/

func hasCycle(graph [][]int, checked []bool, visited []bool, v int) bool {
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

	return false
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	if numCourses <= 1 || prerequisites == nil || len(prerequisites) == 0 {
		return true
	}

	graph := make([][]int, numCourses)
	for i := 0; i < numCourses; i++ {
		graph[i] = make([]int, 0)
	}

	for _, v := range prerequisites {
		idx := v[1]
		val := v[0]
		graph[idx] = append(graph[idx], val)
	}

	checked := make([]bool, numCourses)

	for i := 0; i < numCourses; i++ {
		visited := make([]bool, numCourses)
		if !checked[i] && hasCycle(graph, checked, visited, i) {
			return false
		}
	}

	return true
}

func canFinishTopoSort(numCourses int, prerequisites [][]int) bool {
	if numCourses <= 1 || prerequisites == nil || len(prerequisites) == 0 {
		return true
	}

	graph := make([][]int, numCourses)
	for i := 0; i < numCourses; i++ {
		graph[i] = make([]int, 0)
	}

	inDegrees := make([]int, numCourses)
	for _, v := range prerequisites {
		// set graph
		arr := graph[v[1]]
		arr = append(arr, v[0])
		graph[v[1]] = arr
		// set inDegrees
		inDegrees[v[0]]++
	}

	queue := make([]int, 0)
	for i, v := range inDegrees {
		if v == 0 {
			queue = append(queue, i)
		}
	}

	count := 0
	for len(queue) != 0 {
		// pop
		idx := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		count++

		arr := graph[idx]
		for _, v := range arr {
			inDegrees[v]--
			if inDegrees[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	return count == numCourses
}
