package main

var allPaths [][]int

func allPathsSourceTarget(graph [][]int) [][]int {
	allPaths = make([][]int, 0, len(graph))
	allPathsSourceTargetHelper(graph, 0, []int{})
	return allPaths
}

func allPathsSourceTargetHelper(graph [][]int, node int, path []int) {
	// 向路径添加节点
	path = append(path, node)

	if node == len(graph)-1 {
		allPaths = append(allPaths, path)
		return
	}

	for _, v := range graph[node] {
		allPathsSourceTargetHelper(graph, v, path)
	}

	// 从路径移除节点
	path = path[: len(path)-1 : len(path)-1]
}
