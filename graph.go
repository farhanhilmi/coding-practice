package main

import "fmt"

var graph = map[string][]string{
	"A": {"B", "C"},
	"B": {"D", "E"},
	"C": {"F"},
	"D": {},
	"E": {"F"},
	"F": {},
}

func dfs(graph map[string][]string, start string) {
	// Create a map to keep track of visited vertices
	visited := make(map[string]bool)
	// Create a stack and initialize it with the starting vertex
	stack := []string{start}

	// Loop until the stack is empty
	for len(stack) > 0 {
		// Pop the top vertex from the stack
		vertex := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// If the vertex has not been visited
		if !visited[vertex] {
			// Mark the vertex as visited
			visited[vertex] = true
			// Process the current vertex (in this case, print it)
			fmt.Println(vertex)

			// Add unvisited neighbors to the stack
			for _, neighbor := range graph[vertex] {
				if !visited[neighbor] {
					stack = append(stack, neighbor)
				}
			}
		}
	}
}

// func main() {
// 	dfs(graph, "A")
// }
