#!/bin/python3

import math
import os
import random
import re
import sys
from collections import deque

#
# Complete the 'bfs' function below.
#
# The function is expected to return an INTEGER_ARRAY.
# The function accepts following parameters:
#  1. INTEGER n
#  2. INTEGER m
#  3. 2D_INTEGER_ARRAY edges
#  4. INTEGER s
#

def buildGraph(edges):
    graph = {}
    
    for edge in edges:
        a, b = edge
        graph.setdefault(a, []).append(b)
        graph.setdefault(b, []).append(a)
    
    return graph

def bfs(n, m, edges, s):
    # Write your code here
    graph = buildGraph(edges)

    # Initialize distances to -1 (unreachable)
    distances = [-1] * n
    distances[s - 1] = 0  # Distance to the starting node is 0

    queue = deque([s])

    while queue:
        current_node = queue.popleft()

        for neighbor in graph[current_node]:
            if distances[neighbor - 1] == -1:  # If the neighbor is not visited
                distances[neighbor - 1] = distances[current_node - 1] + 6
                queue.append(neighbor)

    # Exclude the starting node from the distances array
    distances.pop(s - 1)

    return distances

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    q = int(input().strip())

    for q_itr in range(q):
        first_multiple_input = input().rstrip().split()

        n = int(first_multiple_input[0])

        m = int(first_multiple_input[1])

        edges = []

        for _ in range(m):
            edges.append(list(map(int, input().rstrip().split())))

        s = int(input().strip())

        result = bfs(n, m, edges, s)

        fptr.write(' '.join(map(str, result)))
        fptr.write('\n')

    fptr.close()
