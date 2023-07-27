package main

import "fmt"

func main() {
	// Sample input grid
	grid := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}

	// Function to generate
	// New Generation inplace
	NewGeneration(grid)

	// Displaying the grid
	Print(grid)
}

// Print the Grid
func Print(grid [][]int) {
	p := len(grid)
	q := len(grid[0])

	for i := 0; i < p; i++ {
		for j := 0; j < q; j++ {
			fmt.Print(grid[i][j])
		}
		fmt.Println()
	}
}

// Save the grid after new generation
func save(grid [][]int, row int, col int) bool {
	return (len(grid) > row && len(grid[0]) > col && row >= 0 && col >= 0)
}

// Compare cells to find alive neighbours
func NewGeneration(grid [][]int) {
	p := len(grid)
	q := len(grid[0])

	// Possible neighboring
	// indexes
	u := []int{1, -1, 0, 1, -1, 0, 1, -1}
	v := []int{0, 0, -1, -1, -1, 1, 1, 1}

	for i := 0; i < p; i++ {
		for j := 0; j < q; j++ {
			// IF the initial value of the grid(i, j) is 1
			if grid[i][j] > 0 {
				for k := 0; k < 8; k++ {
					if save(grid, i+u[k], j+v[k]) && grid[i+u[k]][j+v[k]] > 0 {
						// If initial value > 0, just increment it by 1
						grid[i][j]++
					}
				}
			} else {
				// IF the initial value of the grid(i, j) is 0
				for k := 0; k < 8; k++ {
					if save(grid, i+u[k], j+v[k]) && grid[i+u[k]][j+v[k]] > 0 {
						// If initial value <= 0 just decrement it by 1
						grid[i][j]--
					}
				}
			}
		}
	}

	// Generating new Generation.
	// Now the magnitude of the grid will represent number of neighbours
	for i := 0; i < p; i++ {
		for j := 0; j < q; j++ {
			// If initial value was 1.
			if grid[i][j] > 0 {
				// Since Any live cell with < 2 live neighbors dies
				if grid[i][j] < 3 {
					grid[i][j] = 0
				} else if grid[i][j] <= 4 {
					// Since Any live cell with 2 or 3 live neighbors live
					grid[i][j] = 1
				} else if grid[i][j] > 4 {
					// Since Any live cell with > 3 live neighbors dies
					grid[i][j] = 0
				}
			} else {
				if grid[i][j] == -3 {
					// Since Any dead cell with exactly 3 live neighbors becomes a live cell
					grid[i][j] = 1
				} else {
					grid[i][j] = 0
				}
			}
		}
	}
}
