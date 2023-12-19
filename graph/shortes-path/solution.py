edges = [
  ['w', 'x'],
  ['x', 'y'],
  ['z', 'y'],
  ['z', 'v'],
  ['w', 'v']
]


def buildGraph(edges):
    graph = {}

    for edge in edges:
        a, b = edge
        graph.setdefault(a, []).append(b)
        graph.setdefault(b, []).append(a)

    return graph


def shortest_path(edges, src, dst):
    graph = buildGraph(edges)
    queue = [[src, 0]]
    visited = set()
    visited.add(src)

    while len(queue) > 0:
        current, distance = queue.pop(0)

        if current == dst:
            return distance

        for neighbor in graph[current]:
            if neighbor not in visited:
                visited.add(neighbor)
                queue.append([neighbor, distance+1])

    return -1

print(shortest_path(edges, 'w', 'z'))