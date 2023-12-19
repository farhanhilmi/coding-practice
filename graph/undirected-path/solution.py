edges = [
  ('i', 'j'),
  ('k', 'i'),
  ('m', 'k'),
  ('k', 'l'),
  ('o', 'n')
]


def buildGraph(edges):
    graph = {}

    for edge in edges:
        a, b = edge
        graph.setdefault(a, []).append(b)
        graph.setdefault(b, []).append(a)

    return graph


def hasPath(graph, src, dst, visited):
    if src in visited:
        return False
    
    if src == dst:
        return True

    visited.add(src)
    for neighbor in graph[src]:
        if hasPath(graph, neighbor, dst, visited):
            return True
        
    return False

def undirectedPath(edges, nodeA, nodeB):
    graph = buildGraph(edges)
    return hasPath(graph, nodeA, nodeB, set())



print(undirectedPath(edges, 'j', 'm'))