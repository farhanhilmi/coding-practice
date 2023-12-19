graph = {
  0: [8, 1, 5],
  1: [0],
  5: [0, 8],
  8: [0, 5],
  2: [3, 4],
  3: [2, 4],
  4: [3, 2]
}


def explore(graph, current, visited):
    if current in visited:
        return False
    
    visited.add(current)
    
    for neighbor in graph[current]:
        explore(graph, neighbor, visited)

    return True

def connectedComponentCount(graph):
    count = 0

    visited = set()
    for node in graph:
        if explore(graph, node, visited):
            count += 1
    
    return count

print(connectedComponentCount(graph))