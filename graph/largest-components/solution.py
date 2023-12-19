graph = {
  0: [8, 1, 5],
  1: [0],
  5: [0, 8],
  8: [0, 5],
  2: [3, 4],
  3: [2, 4],
  4: [3, 2]
}

def explore_size(graph, node, visited):
    if node in visited:
        return 0
    
    visited.add(node)

    size = 1
    for neighbor in graph[node]:
        size += explore_size(graph, neighbor, visited)
        
    return size

def largest_component(graph):
    largest = 0
    for node in graph:
       size = explore_size(graph, node, set())
       if size > largest:
           largest = size

    return largest

print(largest_component(graph))