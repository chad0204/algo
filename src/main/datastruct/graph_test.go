package datastruct

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := buildGraph(numCourses, prerequisites)
	path := make([]bool, numCourses)
	visited := make([]bool, numCourses)
	hasCycle := false
	//因为课程里面可能包含多个隔离的图, 比如1->2->3, 4<->5。所以要从每个节点开始遍历
	for i := 0; i < len(graph); i++ {
		dfs(graph, path, visited, i, &hasCycle)
	}
	return !hasCycle
}

func dfs(graph [][]int, path []bool, visited []bool, s int, hasCycle *bool) {
	if path[s] {
		*hasCycle = true
		return
	}
	//遍历过了
	if visited[s] {
		return
	}
	if *hasCycle {
		return
	}
	visited[s] = true
	path[s] = true
	for _, v := range graph[s] {
		dfs(graph, path, visited, v, hasCycle)
	}
	path[s] = false
}

// 构建邻接表
func buildGraph(numCourses int, prerequisites [][]int) [][]int {
	graph := make([][]int, numCourses)
	for i := 0; i < numCourses; i++ {
		graph[i] = make([]int, 0)
	}
	for _, v := range prerequisites {
		from, to := v[0], v[1]
		graph[from] = append(graph[from], to)
	}
	return graph
}
