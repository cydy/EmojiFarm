package main

import (
	"fmt"
	"math/rand"
)

// Randomly place a 3x3 grid inside a larger grid and pad the surrounding area with edge values.
// Accepts a *rand.Rand instance for deterministic randomness.
func ExpandAndPadGrid(grid3x3 [][]int, targetRows, targetCols, edgeBuffer int, rng *rand.Rand) [][]int {
	// Initialize the larger grid (targetRows x targetCols)
	expandedGrid := make([][]int, targetRows)
	for i := range expandedGrid {
		expandedGrid[i] = make([]int, targetCols)
	}

	// Get the dimensions of the 3x3 grid (which are always 3x3 here)
	smallRows, smallCols := len(grid3x3), len(grid3x3[0])

	// Randomly place the center of the 3x3 grid in the larger grid, considering the edge buffer
	// Use the passed-in rng instance
	// rand.Seed(time.Now().UnixNano()) // Remove local seeding
	centerRow := rng.Intn(targetRows-smallRows-2*edgeBuffer) + edgeBuffer
	centerCol := rng.Intn(targetCols-smallCols-2*edgeBuffer) + edgeBuffer

	// Copy the 3x3 grid into the center of the expanded grid
	for i := range smallRows {
		for j := range smallCols {
			expandedGrid[centerRow+i][centerCol+j] = grid3x3[i][j]
		}
	}

	// Now fill the surrounding area with edge values of the 3x3 grid
	for i := range targetRows {
		for j := range targetCols {
			if expandedGrid[i][j] == 0 { // Empty spot, needs padding
				expandedGrid[i][j] = getNearestEdgeValue(grid3x3, centerRow, centerCol, smallRows, smallCols, i, j)
			}
		}
	}

	return expandedGrid
}

// Get the nearest edge value from the 3x3 grid to fill padding.
func getNearestEdgeValue(grid3x3 [][]int, centerRow, centerCol, smallRows, smallCols, targetRow, targetCol int) int {
	// Determine the nearest edge in the 3x3 grid
	var nearestRow, nearestCol int

	// Determine the nearest row
	if targetRow < centerRow {
		nearestRow = 0
	} else if targetRow >= centerRow+smallRows {
		nearestRow = smallRows - 1
	} else {
		nearestRow = targetRow - centerRow
	}

	// Determine the nearest column
	if targetCol < centerCol {
		nearestCol = 0
	} else if targetCol >= centerCol+smallCols {
		nearestCol = smallCols - 1
	} else {
		nearestCol = targetCol - centerCol
	}

	// Return the value from the nearest edge in the 3x3 grid
	return grid3x3[nearestRow][nearestCol]
}

func test() {
	// Example 3x3 grid
	grid3x3 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	// Define the target size (e.g., 10x13) and edge buffer
	targetRows, targetCols, edgeBuffer := 10, 13, 2

	// For testing, create a local rng instance if needed
	seed := int64(12345) // Example seed for testing
	source := rand.NewSource(seed)
	rng := rand.New(source)

	// Expand and pad the 3x3 grid using the test rng
	expandedGrid := ExpandAndPadGrid(grid3x3, targetRows, targetCols, edgeBuffer, rng)

	// Print the expanded and padded grid
	for _, row := range expandedGrid {
		fmt.Println(row)
	}
}
