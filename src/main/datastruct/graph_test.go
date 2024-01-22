package datastruct

//200. 岛屿数量
func numIslands(grid [][]byte) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}
	row := len(grid)
	col := len(grid[0])
	res := 0
	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			if grid[r][c] == '1' {
				res++
				numIslandsDfs(grid, r, c)
			}
		}
	}
	return res
}
func numIslandsDfs(grid [][]byte, r, c int) {
	//超界
	if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) {
		return
	}
	if grid[r][c] == '0' {
		return
	}
	if grid[r][c] == '2' {
		return
	}
	grid[r][c] = '2'
	numIslandsDfs(grid, r-1, c)
	numIslandsDfs(grid, r+1, c)
	numIslandsDfs(grid, r, c-1)
	numIslandsDfs(grid, r, c+1)
}


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
