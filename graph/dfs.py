graph = {
    'a': ['c','b'],
    'b': ['d'],
    'c': ['e'],
    'd': ['f'],
    'e': [],
    'f': []
}



def depthFirstPrint(graph, source):
    stack = [source]

    while len(stack) > 0:
        current = stack.pop()
        print(current)

        for neighbor in graph[current]:
            stack.append(neighbor)
    
def depthFirstPrintRecursive(graph, source):
    print(source)
    for neighbor in graph[source]:
        depthFirstPrintRecursive(graph, neighbor)

depthFirstPrintRecursive(graph, 'a') # abdfce