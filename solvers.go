package main

import (
	"fmt"
)

func naive(width int, height int, painted [][]bool) []string {
	var ops []string
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if painted[y][x] {
				ops = append(ops, fmt.Sprintf("FILL,%d,%d,1", x, y))
			}
		}
	}
	return ops
}

func greedy(width int, height int, painted [][]bool) []string {
	visited := make([][]bool, height)
	for i := 0; i < height; i++ {
		visited[i] = make([]bool, width)
	}

	var ops []string
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if visited[y][x] {
				continue
			} else if painted[y][x] {
				size := 1
				isNeighborWhite := false
				for y + size < height && x + size < width {
					for k := 0; k < size + 1; k++ {
						if !painted[y+k][x+size] {
							isNeighborWhite = true
						}
						if isNeighborWhite {
							break
						}
					}
					if isNeighborWhite {
						break
					}
					for k := size - 1; k >= 0; k-- {
						if !painted[y+size][x+k] {
							isNeighborWhite = true
						}
						if isNeighborWhite {
							break
						}
					}
					if isNeighborWhite {
						break
					} else {
						for k := 0; k < size + 1; k++ {
							visited[y+k][x+size] = true
						}
						for k := size - 1; k >= 0; k-- {
							visited[y+size][x+k] = true
						}
						size++
					}
				}
				ops = append(ops, fmt.Sprintf("FILL,%d,%d,%d", x, y, size))
			}
		}
	}
	return ops
}
