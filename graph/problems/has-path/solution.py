graph = {
  'f': ['g', 'i'],
  'g': ['h'],
  'h': [],
  'i': ['g', 'k'],
  'j': ['i'],
  'k': []
}

def has_path(graph, src, dst): # BFS
  if src == dst:
    return True
  
  queue = [src]
  
  while len(queue) > 0:
    current = queue.pop(0)
    if current == dst:
        return True
    
    for neighbor in graph[current]:
      queue.append(neighbor)
      

  return False

def has_path_recursive(graph, src, dst): # DFS with recursive
  if src == dst:
    return True
  
  for neigbor in graph[src]:
    if has_path_recursive(graph, neigbor, dst) == True:
      return True
      
  return False

print(has_path(graph, 'f', 'k'))