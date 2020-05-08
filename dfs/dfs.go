package main

import "fmt"

func addEdge(a, b int, adj [][]int) {
	adj[a] = append(adj[a], b)
	adj[b] = append(adj[b], a)
}

func dfs(node int, adj [][]int, visited []int) {
	visited[node] = 1
	fmt.Printf("%d -> ", node)
	for _, child := range adj[node] {
		if visited[child] == 0 {
			dfs(child, adj, visited)
		}
	}
}

func main() {
	numberOfNodes := 7

	visited := make([]int, numberOfNodes)

	adj := make([][]int, numberOfNodes)
	for i := range adj {
		adj[i] = make([]int, 0)
	}

	addEdge(1, 2, adj)
	addEdge(2, 3, adj)
	addEdge(2, 4, adj)
	addEdge(4, 5, adj)
	addEdge(4, 6, adj)

	dfs(1, adj, visited) // Prints 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 
}
