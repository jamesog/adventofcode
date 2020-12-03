package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Take a grid and return the number of trees encountered
func travel(grid []string, right, down int) int {
	// For each move, go right 3, down 1, record if tree encountered
	var trees int
	var x, y int
	for {
		x += right
		y += down
		if y >= len(grid) {
			// We reached the bottom
			break
		}
		if x > len(grid[y]) {
			// We reached the edge of the grid without reaching the bottom,
			// so expand the grid to the right
			grow(grid)
		}
		sq := grid[y][x]
		if sq == '#' {
			// fmt.Printf("(%d, %d) TREE!\n", x, y)
			trees++
		}
	}
	return trees
}

// Expand the grid by repeating each row
func grow(grid []string) {
	for row := range grid {
		grid[row] = grid[row] + grid[row]
	}
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Couldn't open input: %v", err)
	}

	var grid []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}

	trees := travel(grid, 3, 1)
	fmt.Println("Part 1:", trees)

	trees *= travel(grid, 1, 1)
	trees *= travel(grid, 5, 1)
	trees *= travel(grid, 7, 1)
	trees *= travel(grid, 1, 2)
	fmt.Println("Part 2:", trees)
}
