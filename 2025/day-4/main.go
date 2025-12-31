package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(data), "\n")

	fmt.Println("part 1", part1(lines))
	fmt.Println("part 2", part2(lines))
}

func part1(lines []string) int {
	k := 4
	accessible := 0
	grid := make([][]byte, 0)

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		row := make([]byte, 0)
		for j := 0; j < len(line); j++ {
			row = append(row, line[j])
		}
		grid = append(grid, row)
	}

	for i := 0; i < len(grid); i++ {
		row := grid[i]
		for j := 0; j < len(row); j++ {
			count := countPaperNeighbors(grid, i, j)
			if count < k && row[j] == '@' {
				accessible++
			}
		}
	}

	return accessible
}

func countPaperNeighbors(grid [][]byte, x, y int) int {
	count := 0
	neighbors := [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

	for _, n := range neighbors {
		i := x + n[0]
		j := y + n[1]

		if i >= 0 && j >= 0 && i < len(grid) && j < len(grid[i]) && grid[i][j] == '@' {
			count++
		}
	}

	return count
}

func part2(lines []string) int {
	k := 4
	removed := 0
	grid := make([][]byte, 0)

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		row := make([]byte, 0)
		for j := 0; j < len(line); j++ {
			row = append(row, line[j])
		}
		grid = append(grid, row)
	}

	for newRemoved := removeAccessible(grid, k); newRemoved > 0; {
		removed += newRemoved
		newRemoved = removeAccessible(grid, k)
	}

	return removed
}

func removeAccessible(grid [][]byte, k int) int {
	removed := 0

	for i := 0; i < len(grid); i++ {
		row := grid[i]
		for j := 0; j < len(row); j++ {
			count := countPaperNeighbors(grid, i, j)
			if count < k && row[j] == '@' {
				removed++
				grid[i][j] = '.'
			}
		}
	}

	return removed
}
