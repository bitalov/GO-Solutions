// https://leetcode.com/problems/is-graph-bipartite


func isOK(node, color int, painted []int, graph [][]int) bool {
	if painted[node] != 0 {
		return painted[node] == color
	}

	painted[node] = color

	for _, i := range graph[node] {
		if !isOK(i, -color, painted, graph) {
			return false
		}
	}

	return true
}

func isBipartite(graph [][]int) bool {
	n := len(graph)
	painted := make([]int, n)

	for i := 0; i < n; i++ {
		if painted[i] == 0 && !isOK(i, 1, painted, graph) {
			return false
		}
	}

	return true
}
