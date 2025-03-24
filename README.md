# optimal_art
Solver for the "Optimal art" problem of [primers.xyz](https://primers.xyz/0).

## Solutions
### Naive approach
Algorithm for the naive approach:
```
parse input -> a [][]bool painted such that painted[i][j] is true if pixel (i, j) is black.
[]string operations
for i in painted height
    for j in painted width
        if painted[i][j] is true then
            append "FILL,j,i,1" to back of operations list
write each element of operations to file
```

This approach produces a file too big to be accepted by primers website now.

Score (obtained with `wc -l`): 84784

### Greedy approach
Fill the biggest square you can
```
parse input -> a [][]bool painted such that painted[i][j] is true if pixel (i, j) is black.
initialize a visited [][]bool of same size than painted such that visited[i][j] is true if painted[i][j] was visited
for i in painted height
    for j in painted width
        if visited[i][j] then do nothing
        else if painted[i][j] then
            size = 1
            isNeighborWhite = false
            while (i+size < height && j+size < width)
                for k = 0; k < size+1; k++
                    if !painted[i+k, j+size]
                        isNeighborWhite = true
                    if isNeighborWhite then break
                if isNeighborWhite then break
                for k = size - 1; k >= 0; k--
                    if !painted[i+size, j+k]
                        isNeighborWhite = true
                    if isNeighborWhite then break
                if isNeighborWhite then break
                else
                    for k = 0; k < size+1; k++
                        visited[i+k, j+size] = true
                    for k = size - 1; k >= 0; k--
                        visited[i+size, j+k] = true
                    incremente size
            fill, j, i, size
```

Score: 4577, a good improvement

## Other ideas
Other ideas that were not implemented:

- count black and white cell, if black > white then fill only else one fill then
only erase
- dynamic programming approach, solve pb for painting of incrementing size
