package main

import (
	"fmt"
	"hash/fnv"
	"math/rand"
	"sort"
)

// VerboseLogger is a helper for verbose logging
type VerboseLogger struct {
	enabled bool
}

// NewVerboseLogger creates a new verbose logger
func NewVerboseLogger(enabled bool) *VerboseLogger {
	return &VerboseLogger{enabled: enabled}
}

// Log prints a message if verbose logging is enabled
func (v *VerboseLogger) Log(format string, args ...interface{}) {
	if v.enabled {
		fmt.Printf(format, args...)
	}
}

// Logln prints a message with newline if verbose logging is enabled
func (v *VerboseLogger) Logln(args ...interface{}) {
	if v.enabled {
		fmt.Println(args...)
	}
}

// EmojiGrid structure similar to Python class
type EmojiGrid struct {
	rows, cols    int
	canvasInt     [][]int
	canvasStrings [][]string
	edgeBuffer    int
	sectionNums   []int
	barrierNums   []int
	seed          int64
	rng           *rand.Rand // Add local random number generator
}

// NewEmojiGrid initializes a new grid
func NewEmojiGrid(seed int64) *EmojiGrid {
	rows, cols := 10, 13
	// Create a new random source and generator for this grid instance
	source := rand.NewSource(seed)
	rng := rand.New(source)

	canvasInt := [][]int{
		{1, 2, 3},
		{8, 9, 4},
		{7, 6, 5},
	}

	return &EmojiGrid{
		rows:       rows,
		cols:       cols,
		canvasInt:  canvasInt,
		edgeBuffer: 3,
		seed:       seed,
		rng:        rng, // Initialize the local rng
	}
}

// ReplaceSection replaces one number with another
func (e *EmojiGrid) ReplaceSection(cur, new int) {
	for i := range e.canvasInt {
		for j := range e.canvasInt[i] {
			if e.canvasInt[i][j] == cur {
				e.canvasInt[i][j] = new
			}
		}
	}
}

func (e *EmojiGrid) GetBarrierNums() []int {
	var temp []int
	for _, x := range e.getSections() {
		if x%2 == 0 {
			temp = append(temp, x)
		}
	}
	sort.Ints(temp)
	return temp
}

// SplitGen generates a split based on random values (placeholder logic)
func (e *EmojiGrid) SplitGen(logger *VerboseLogger) {
	// Possible direction: 0 = none, 1 = keep '2' or '4' barrier, 2 = keep '6' or '8' barrier, 3 = keep both
	dirX := e.rng.Intn(4)
	logger.Log("Direction X: %d\n", dirX)
	dirY := e.rng.Intn(4)
	logger.Log("Direction Y: %d\n", dirY)

	// Conditions to skip the split
	if (dirX == 0 || dirY == 0) && ((dirX == 1 || dirY == 1) || (dirX == 2 || dirY == 2)) {
		logger.Logln("No split")
		if e.rng.Intn(2) == 0 { // Retry split condition
			dirX = e.rng.Intn(4)
			dirY = e.rng.Intn(4)
			logger.Logln("Retry split!")
			return
		} else {
			for i := range e.canvasInt {
				for j := range e.canvasInt[i] {
					if e.canvasInt[i][j] > 0 {
						e.canvasInt[i][j] = 1
					}
				}
			}
			logger.Logln("No split again!")
			return
		}
	}

	// Clockwise order from NW (implementing based on dirX and dirY)
	if dirY != 1 && dirY != 3 {
		if e.rng.Intn(2) == 0 {
			e.ReplaceSection(2, 1)
			e.ReplaceSection(3, 1)
		} else {
			e.ReplaceSection(2, 2+e.rng.Intn(2)*2-1) // 1 or 3
		}
	}

	if dirX != 2 && dirX != 3 {
		if e.rng.Intn(2) == 0 {
			e.ReplaceSection(4, 3)
			e.ReplaceSection(5, 3)
		} else {
			e.ReplaceSection(4, 4+e.rng.Intn(2)*2-1) // 3 or 5
		}
	}

	if dirY != 2 && dirY != 3 {
		if e.rng.Intn(2) == 0 {
			e.ReplaceSection(6, 5)
			e.ReplaceSection(7, 5)
		} else {
			e.ReplaceSection(6, 6+e.rng.Intn(2)*2-1) // 5 or 7
		}
	}

	if dirX != 1 && dirX != 3 {
		if e.rng.Intn(2) == 0 {
			e.ReplaceSection(8, 7)
			e.ReplaceSection(1, 7)
		} else {
			e.ReplaceSection(8, e.rng.Intn(2)*6+1) // 1 or 7
		}
	}

	logger.Logln("Post splitGen:")
	e.PrintCanvas(logger)
}

// findValidNeighbors finds non-barrier neighbors for a given section number
func (e *EmojiGrid) findValidNeighbors(sectionNum int) []int {
	neighbors := make(map[int]bool)
	rows := len(e.canvasInt)
	cols := len(e.canvasInt[0]) // Assuming a rectangular grid

	for r := range rows {
		for c := range cols {
			if e.canvasInt[r][c] == sectionNum {
				// Check adjacent cells (up, down, left, right)
				adj := [][]int{{r - 1, c}, {r + 1, c}, {r, c - 1}, {r, c + 1}}
				for _, pos := range adj {
					nr, nc := pos[0], pos[1]
					// Check bounds
					if nr >= 0 && nr < rows && nc >= 0 && nc < cols {
						neighborVal := e.canvasInt[nr][nc]
						// Add if it's a non-barrier (odd), not the section itself, AND not 9.
						if neighborVal != sectionNum && neighborVal%2 != 0 && neighborVal != 9 {
							neighbors[neighborVal] = true
						}
					}
				}
			}
		}
	}

	// Convert map keys to slice
	var result []int
	for n := range neighbors {
		result = append(result, n)
	}
	return result
}

// DecideSplits decides which splits to keep
func (e *EmojiGrid) DecideSplits(logger *VerboseLogger) {
	// decide to keep split or remove & merge

	// If only one section exists, return
	if len(e.getSections()) == 1 {
		logger.Logln("Only one section exists, skipping splits.")
		return
	}

	dirs := e.GetBarrierNums() // Get current barriers (even numbers)
	activeDirs := make([]int, len(dirs))
	copy(activeDirs, dirs) // Create a copy to iterate over, as dirs might change
	sort.Ints(activeDirs)  // Sort activeDirs for deterministic iteration order

	// Iterate over remaining splits
	for _, i := range activeDirs {
		// Check if the barrier 'i' still exists (wasn't merged already)
		stillExists := false
		for _, currentSection := range e.getSections() {
			if currentSection == i {
				stillExists = true
				break
			}
		}
		if !stillExists {
			continue // Skip if this barrier was already merged
		}

		if e.rng.Intn(2) == 0 { // Remove split
			logger.Log("Attempting to remove barrier: %d\n", i)
			validNeighbors := e.findValidNeighbors(i)
			logger.Log("Valid neighbors for %d: %v\n", i, validNeighbors)

			if len(validNeighbors) > 0 {
				// Choose a random valid neighbor to merge into
				// Sort validNeighbors before choosing randomly to ensure determinism
				sort.Ints(validNeighbors)
				mergeTarget := validNeighbors[e.rng.Intn(len(validNeighbors))]
				logger.Log("Merging %d into %d\n", i, mergeTarget)
				e.ReplaceSection(i, mergeTarget)

				// Remove the merged barrier from the original dirs slice
				for idx, val := range dirs {
					if val == i {
						dirs = append(dirs[:idx], dirs[idx+1:]...)
						break
					}
				}
			} else {
				logger.Log("Could not find valid neighbors for %d, keeping barrier.\n", i)
				// Keep split (do nothing) or implement fallback
			}
		} else {
			logger.Log("Keeping barrier: %d\n", i)
			// Keep split
			continue
		}
	}

	// DECIDE ONE SPLIT

	logger.Log("Remaining barriers: %v\n", dirs) // Print barriers that were kept
	logger.Logln("Post decideSplits:")
	e.PrintCanvas(logger)
}

// Helper function to get coordinates for each section
func (e *EmojiGrid) getSectionCoords() map[int][][2]int {
	coordsMap := make(map[int][][2]int)
	for r, row := range e.canvasInt {
		for c, val := range row {
			// Collect coordinates for sections 1, 3, 5, 7, and 9
			if val%2 != 0 || val == 9 { // Odd numbers OR 9
				coordsMap[val] = append(coordsMap[val], [2]int{r, c})
			}
		}
	}
	return coordsMap
}

// Helper function to check if 4 coordinates form a 2x2 square
func is2x2Square(coords [][2]int) bool {
	if len(coords) != 4 {
		return false // Must have exactly 4 coordinates
	}

	minRow, maxRow := coords[0][0], coords[0][0]
	minCol, maxCol := coords[0][1], coords[0][1]

	for i := 1; i < 4; i++ {
		r, c := coords[i][0], coords[i][1]
		if r < minRow {
			minRow = r
		}
		if r > maxRow {
			maxRow = r
		}
		if c < minCol {
			minCol = c
		}
		if c > maxCol {
			maxCol = c
		}
	}

	return (maxRow-minRow == 1) && (maxCol-minCol == 1)
}

func (e *EmojiGrid) DecideBarriers(logger *VerboseLogger) {
	// Decide whether to keep remaining split as a barrier
	if e.rng.Intn(3) != 0 {
		// Replace all barriers (evens) with 2
		logger.Log("Replace all barriers (evens) with 2\n")
		for i := 2; i < 10; i += 2 {
			e.ReplaceSection(i, 2)
		}
	}

	// --- Handle Section 9 Replacement ---
	// Get section coordinates (includes counts via length)
	sectionCoords := e.getSectionCoords()

	// Check if section 9 exists
	coords9, hasNine := sectionCoords[9]
	// It's assumed section 9 should only have 1 coordinate at this stage for the logic to work
	if hasNine && len(coords9) != 1 {
		logger.Log("Warning: Section 9 has %d coordinates, expected 1. Geometric check might be unreliable.\n", len(coords9))
		// Proceed anyway, but be aware
	}

	if hasNine {
		logger.Log("Handling section 9 replacement with order: Rule 2 -> Rule 1 -> Rule 3\n")
		mergeTarget := -1 // Initialize merge target

		// --- Start Reordered Logic ---

		// 1. Apply Rule 2: Frequency 3 with Geometric Tie-breaker
		logger.Log("Checking Rule 2 (Freq 3 + Geometry)...\n")
		sectionsWithFreq3 := []int{}
		for section, coords := range sectionCoords {
			// Only consider main sections (1, 3, 5, 7) for frequency check
			if section != 9 && section%2 != 0 {
				if len(coords) == 3 {
					sectionsWithFreq3 = append(sectionsWithFreq3, section)
				}
			}
		}
		logger.Log("Sections with frequency 3: %v\n", sectionsWithFreq3)

		if len(sectionsWithFreq3) == 1 {
			// Only one section has frequency 3 - simple case
			mergeTarget = sectionsWithFreq3[0]
			logger.Log("Rule 2 Applied (Unique Freq 3): Merging 9 into section %d\n", mergeTarget)
		} else if len(sectionsWithFreq3) > 1 {
			sort.Ints(sectionsWithFreq3) // Sort for deterministic iteration
			// Tie-breaker: Check which one forms a 2x2 square with 9
			logger.Log("Rule 2 Tie-breaker: Checking for 2x2 square formation...\n")
			formingSquareSections := []int{}
			coords9 = sectionCoords[9] // Ensure we have the coords for 9

			if len(coords9) == 1 { // Geometric check requires single coord for 9
				for _, section := range sectionsWithFreq3 {
					coordsSection := sectionCoords[section]
					combinedCoords := append([][2]int{}, coordsSection...) // Create a copy
					combinedCoords = append(combinedCoords, coords9[0])    // Add coord of 9

					if is2x2Square(combinedCoords) {
						logger.Log("  - Section %d forms 2x2 square with 9.\n", section)
						formingSquareSections = append(formingSquareSections, section)
					} else {
						logger.Log("  - Section %d does NOT form 2x2 square with 9.\n", section)
					}
				}

				if len(formingSquareSections) == 1 {
					// Exactly one section forms a 2x2 square - perfect!
					mergeTarget = formingSquareSections[0]
					logger.Log("Rule 2 Applied (Geometric Tie-breaker): Merging 9 into section %d\n", mergeTarget)
				} else if len(formingSquareSections) > 1 {
					sort.Ints(formingSquareSections)       // Sort potential targets
					mergeTarget = formingSquareSections[0] // Choose deterministically
					logger.Log("Rule 2 Applied (Geometric Tie-breaker - Multiple): Merging 9 into section %d (smallest forming square)\n", mergeTarget)
				} else {
					logger.Log("Rule 2 Tie-breaker Failed: No section formed a 2x2 square.\n")
				}
			} else {
				logger.Log("Rule 2 Tie-breaker Skipped: Section 9 does not have exactly 1 coordinate.\n")
			}
		} else {
			logger.Log("Rule 2 Did Not Apply: No sections with frequency 3.\n")
		}

		// 2. Apply Rule 1: Largest Barrier (only if Rule 2 didn't apply)
		if mergeTarget == -1 {
			logger.Log("Rule 2 did not determine target. Checking Rule 1 (Barriers)...\n")
			remainingBarriers := e.GetBarrierNums()
			if len(remainingBarriers) > 0 {
				// Find and set the barrier with the highest number
				sort.Ints(remainingBarriers)                                  // Sort barriers
				biggestBarrier := remainingBarriers[len(remainingBarriers)-1] // Get largest after sorting
				mergeTarget = biggestBarrier
				logger.Log("Rule 1 Applied: Merging 9 into largest barrier: %d\n", biggestBarrier)
			} else {
				logger.Log("Rule 1 did not apply (no barriers).\n")
			}
		}

		// 3. Apply Rule 3: Highest Frequency (only if Rule 2 and Rule 1 didn't apply)
		if mergeTarget == -1 {
			logger.Log("Rule 1 did not determine target. Checking Rule 3 (Highest Frequency)...\n")

			maxFreq := 0
			sectionsWithMaxFreq := []int{}
			for section, coords := range sectionCoords {
				// Only consider main sections (1, 3, 5, 7)
				if section != 9 && section%2 != 0 {
					count := len(coords)
					if count > maxFreq {
						maxFreq = count
						sectionsWithMaxFreq = []int{section} // Reset
					} else if count == maxFreq {
						sectionsWithMaxFreq = append(sectionsWithMaxFreq, section) // Add to tie
					}
				}
			}
			logger.Log("Highest frequency found: %d, Sections: %v\n", maxFreq, sectionsWithMaxFreq)

			if len(sectionsWithMaxFreq) == 1 {
				mergeTarget = sectionsWithMaxFreq[0]
				logger.Log("Rule 3 Applied: Merging 9 into section %d (highest frequency: %d)\n", mergeTarget, maxFreq)
			} else if len(sectionsWithMaxFreq) > 1 {
				sort.Ints(sectionsWithMaxFreq) // Tie-breaker (smallest number)
				mergeTarget = sectionsWithMaxFreq[0]
				logger.Log("Rule 3 Applied (Tie): Merging 9 into section %d (smallest of max freq sections)\n", mergeTarget)
			} else {
				logger.Log("Rule 3 Did Not Apply: No valid sections found (maxFreq=0?).\n")
			}
		}

		// --- End Reordered Logic ---

		// Perform the replacement if a target was determined
		if mergeTarget != -1 {
			e.ReplaceSection(9, mergeTarget)
		} else {
			// Fallback if NO rules could determine a target
			logger.Log("Warning: Could not determine merge target for section 9 based on rules 2, 1, or 3. Replacing with 1 as fallback.\n")
			e.ReplaceSection(9, 1) // Fallback to section 1
		}
	} else {
		logger.Log("Section 9 not present.\n")
	}
	// --- End Handle Section 9 Replacement ---

	logger.Logln("Post decideBarriers:")
	e.PrintCanvas(logger)
}

// GetSections returns unique numbers in the canvas, sorted
func (e *EmojiGrid) getSections() []int {
	unique := make(map[int]bool)
	for _, row := range e.canvasInt {
		for _, val := range row {
			unique[val] = true
		}
	}
	var sections []int
	for key := range unique {
		sections = append(sections, key)
	}
	sort.Ints(sections) // Sort the sections for deterministic order
	return sections
}

// BuildFarm builds the farm structure
func (e *EmojiGrid) BuildFarm(logger *VerboseLogger) {
	logger.Logln("Building Farm:")
	logger.Logln("-------------")

	logger.Logln("\n1. Generating Split:")
	e.SplitGen(logger)

	logger.Logln("\n2. Deciding Splits:")
	e.DecideSplits(logger)

	logger.Logln("\n3. Deciding Barriers:")
	e.DecideBarriers(logger)

	// Use ExpandAndPadGrid to expand the 3x3 canvas to the full size
	// Pass the grid's rng instance
	expandedGrid := ExpandAndPadGrid(e.canvasInt, e.rows, e.cols, e.edgeBuffer-1, e.rng)
	e.canvasInt = expandedGrid

	e.sectionNums = append(e.getSections(), e.GetBarrierNums()...)
	sort.Ints(e.sectionNums) // Sort combined section numbers before iteration
	e.barrierNums = e.GetBarrierNums()

	e.ConvertCanvasToStrings()

	logger.Logln("\n4. Farm Structure:")
	logger.Log("\tBarriers: %v\n", e.GetBarrierNums())
	logger.Log("\tSections: %v\n", e.getSections())

	logger.Logln("\n5. Building Biomes:")
	for _, sectNum := range e.sectionNums {
		biome := NewBiome(e, e.canvasStrings, sectNum, e.rng)
		biome.Build(logger)
		biome.Replace()
	}
	logger.Logln("\n6. Numeric canvas:")
	e.PrintCanvas(logger)

	logger.Logln("\n7. Emoji canvas:")
	e.PrintFinalEmojis(true)
}

// PrintCanvas prints the canvas (grid)
func (e *EmojiGrid) PrintCanvas(logger *VerboseLogger) {
	for _, row := range e.canvasInt {
		logger.Logln(row)
	}
}

// ConvertCanvasToStrings converts the int canvas to a string canvas
func (e *EmojiGrid) ConvertCanvasToStrings() {
	e.canvasStrings = make([][]string, len(e.canvasInt))
	for i, row := range e.canvasInt {
		e.canvasStrings[i] = make([]string, len(row))
		for j, val := range row {
			e.canvasStrings[i][j] = fmt.Sprintf("%d", val)
		}
	}
}

// PrintFinalEmojis returns the grid as a string and optionally prints it
func (e *EmojiGrid) PrintFinalEmojis(printToConsole bool) string {
	var result string
	for _, row := range e.canvasStrings {
		for _, emoji := range row {
			result += emoji
		}
		result += "\n"
	}
	if printToConsole {
		fmt.Print(result)
	}
	return result
}

// Convert string to a numeric seed using hash function
func hashString(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}
