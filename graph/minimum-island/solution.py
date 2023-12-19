grid = [
  ['W', 'L', 'W', 'W', 'W'],
  ['W', 'L', 'W', 'W', 'W'],
  ['W', 'W', 'W', 'L', 'W'],
  ['W', 'W', 'L', 'L', 'W'],
  ['L', 'W', 'W', 'L', 'L'],
  ['L', 'L', 'W', 'W', 'W'],
]

def explore(grid, row, col, visited):
    row_inbound = 0 <= row and row < len(grid)
    col_inbound = 0 <= col and col < len(grid[0])

    if not row_inbound or not col_inbound:
        return 0
    
    if grid[row][col] == "W":
        return 0
    
    pos = (row, col)
    if pos in visited:
        return 0
    
    visited.add(pos)
    size = 1

    size += explore(grid, row - 1, col, visited)
    size += explore(grid, row + 1, col, visited)
    size += explore(grid, row, col - 1, visited)
    size += explore(grid, row, col + 1, visited)

    return size 

def minimum_island(grid):
    count = float("inf")
    print("count", count)
    visited = set()

    for row in range(len(grid)):
        for col in range(len(grid[0])):
            size = explore(grid, row, col, visited)
            if size > 0 and size < count:
                count = size

    return count

print(minimum_island(grid))