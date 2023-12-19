grid = [
['W', 'L', 'W', 'W', 'W'],
['W', 'L', 'W', 'W', 'W'],
['W', 'W', 'W', 'L', 'W'],
['W', 'W', 'L', 'L', 'W'],
['L', 'W', 'W', 'L', 'L'],
['L', 'L', 'W', 'W', 'W'],
]

def explore(grid, row, col, visited):
    row_inbounds = 0 <= row < len(grid)
    col_inbounds = 0 <= col < len(grid[0])
    if not row_inbounds or not col_inbounds:
      return False
    
    if grid[row][col] == 'W':
        return False

    pos = (row, col)
    if pos in visited:
      return False
    
    visited.add(pos)

    explore(grid, row - 1, col, visited) # UP
    explore(grid, row + 1, col, visited) # DOWN
    explore(grid, row, col - 1, visited) # LEFT
    explore(grid, row, col + 1, visited) # RIGHT

    return True


def island_count(grid):
    count = 0
    visited = set()

    for row in range(len(grid)):
      for col in range(len(grid[0])):
        if explore(grid, row, col, visited) == True:
          count += 1

    return count

print(island_count(grid))