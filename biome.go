//go:build js && wasm

package main

import (
	"fmt"
	"math"
	"math/rand"
)

// Biome represents a section of the farm
type Biome struct {
	grid           *EmojiGrid
	sectionCanvas  [][]string
	sectNum        int
	cords          [][]int
	availableTypes map[int]func(*VerboseLogger)
	rng            *rand.Rand
}

// Helper function to get section coordinates
func getSectionCords(inputCanvas [][]int, sectNum int) [][]int {
	var cords [][]int
	for i, row := range inputCanvas {
		for j, val := range row {
			if val == sectNum {
				cords = append(cords, []int{i, j})
			}
		}
	}
	return cords
}

// Helper function to split numbers using the provided rng
func splitNum(parts, total int, rng *rand.Rand) []int {
	result := make([]int, parts)
	randomIndex := rng.Intn(parts)
	for range total {
		result[randomIndex]++
	}
	return result
}

func NewBiome(grid *EmojiGrid, sectionCanvas [][]string, sectNum int, rng *rand.Rand) *Biome {
	b := &Biome{
		grid:          grid,
		sectionCanvas: sectionCanvas,
		sectNum:       sectNum,
		cords:         getSectionCords(grid.canvasInt, sectNum),
		rng:           rng,
	}

	b.availableTypes = map[int]func(*VerboseLogger){
		0: b.crops,
		1: b.animals,
		2: b.field,
		3: b.pond,
	}

	return b
}

func (b *Biome) replaceSection(gridInput [][]string, cur, new string) {
	for i := range gridInput {
		for j := range gridInput[i] {
			if gridInput[i][j] == cur {
				gridInput[i][j] = new
			}
		}
	}
}

func (b *Biome) crops(logger *VerboseLogger) {
	logger.Log("crops: whole\n")
	crop1 := PLANTS_CROPS[b.rng.Intn(len(PLANTS_CROPS))]
	logger.Log("crop: %s\n", crop1)
	b.replaceSection(b.sectionCanvas, fmt.Sprint(b.sectNum), crop1)
}

func (b *Biome) animals(logger *VerboseLogger) {
	placementCords := b.cords
	numberOfCords := len(placementCords)

	// Fill the entire section with grass background
	for _, cord := range placementCords {
		b.sectionCanvas[cord[0]][cord[1]] = "🌱"
	}

	fillPercentage := b.rng.Float64()*(0.6-0.2) + 0.2
	numCordsToFill := int(math.Round(float64(numberOfCords) * fillPercentage))

	logger.Log("filling %d, %.2f, %d\n", numberOfCords, fillPercentage, numCordsToFill)

	// Shuffle the placement coordinates using b.rng
	b.rng.Shuffle(len(placementCords), func(i, j int) {
		placementCords[i], placementCords[j] = placementCords[j], placementCords[i]
	})

	var animalSet []string
	if b.rng.Intn(2) == 1 {
		logger.Log("animals: livestock\n")
		animalSet = ANIMALS_FARM
	} else {
		logger.Log("animals: birds\n")
		animalSet = ANIMALS_BIRDS
	}

	if b.rng.Intn(4) == 0 { // 1/4 chance one animal biome
		logger.Log("animals #: 1\n")
		animal1 := animalSet[b.rng.Intn(len(animalSet))]
		logger.Log("animal1: %s\n", animal1)
		for i := range numCordsToFill {
			cord := placementCords[i]
			b.sectionCanvas[cord[0]][cord[1]] = animal1
		}
	} else {
		animalMax := min(5, numCordsToFill)
		animalNum := b.rng.Intn(animalMax) + 1
		logger.Log("animals #: %d\n", animalNum)

		animalChoices := randomSample(animalSet, animalNum, b.rng)
		animalDistribution := splitNum(animalNum, numCordsToFill, b.rng)

		z := 0
		for i, animalOutput := range animalChoices {
			for range animalDistribution[i] {
				cord := placementCords[z]
				b.sectionCanvas[cord[0]][cord[1]] = animalOutput
				z++
			}
		}
	}
}

func (b *Biome) field(logger *VerboseLogger) {
	placementCords := b.cords
	numberOfCords := len(placementCords)

	// Fill the entire section with grass background
	for _, cord := range placementCords {
		b.sectionCanvas[cord[0]][cord[1]] = "🌱"
	}

	fillPercentage := b.rng.Float64()*(0.8-0.2) + 0.2
	numCordsToFill := int(math.Round(float64(numberOfCords) * fillPercentage))

	logger.Log("filling %d, %.2f, %d\n", numberOfCords, fillPercentage, numCordsToFill)

	// Shuffle the placement coordinates using b.rng
	b.rng.Shuffle(len(placementCords), func(i, j int) {
		placementCords[i], placementCords[j] = placementCords[j], placementCords[i]
	})

	var plantSet []string
	if b.rng.Intn(2) == 1 {
		logger.Log("plants: field\n")
		plantSet = PLANTS_FIELD
	} else {
		logger.Log("plants: flowers\n")
		plantSet = PLANTS_FLOWERS
	}

	if b.rng.Intn(4) == 0 { // 1/4 chance one plant biome
		logger.Log("plants #: 1\n")
		plant1 := plantSet[b.rng.Intn(len(plantSet))]
		logger.Log("plant1: %s\n", plant1)
		for i := range numCordsToFill {
			cord := placementCords[i]
			b.sectionCanvas[cord[0]][cord[1]] = plant1
		}
	} else {
		plantMax := min(5, numCordsToFill)
		plantNum := b.rng.Intn(plantMax) + 1
		logger.Log("plants #: %d\n", plantNum)

		plantChoices := randomSample(plantSet, plantNum, b.rng)
		plantDistribution := splitNum(plantNum, numCordsToFill, b.rng)

		z := 0
		for i, plantOutput := range plantChoices {
			for range plantDistribution[i] {
				cord := placementCords[z]
				b.sectionCanvas[cord[0]][cord[1]] = plantOutput
				z++
			}
		}
	}
}

// pond builds a pond biome
func (b *Biome) pond(logger *VerboseLogger) {
	logger.Log("Building pond biome\n")
	placementCords := b.cords
	numberOfCords := len(placementCords)
	fillPercentage := b.rng.Float64()*0.4 + 0.1
	numCordsToFill := int(math.Round(float64(numberOfCords) * fillPercentage))
	logger.Log("filling %d, %.2f, %d\n", numberOfCords, fillPercentage, numCordsToFill)

	// Shuffle placementCords using b.rng
	b.rng.Shuffle(len(placementCords), func(i, j int) {
		placementCords[i], placementCords[j] = placementCords[j], placementCords[i]
	})
	placementCords = placementCords[:numCordsToFill]

	b.replaceSection(b.sectionCanvas, fmt.Sprintf("%d", b.sectNum), "🌊")

	pawnSet := POND

	if b.rng.Intn(4) == 0 { // 3/4 chance one pawn biome
		logger.Log("pawns #: 1\n")
		pawn1 := pawnSet[b.rng.Intn(len(pawnSet))]
		logger.Log("pawn1: %s\n", pawn1)
		for _, cord := range placementCords {
			b.sectionCanvas[cord[0]][cord[1]] = pawn1
		}
	} else {
		pawnMax := min(5, numCordsToFill)
		pawnNum := b.rng.Intn(pawnMax) + 1
		logger.Log("pawns #: %d\n", pawnNum)

		pawnChoices := randomSample(pawnSet, pawnNum, b.rng)
		pawnDistribution := splitNum(pawnNum, numCordsToFill, b.rng)

		z := 0
		for i, pawnOutput := range pawnChoices {
			for range pawnDistribution[i] {
				b.sectionCanvas[placementCords[z][0]][placementCords[z][1]] = pawnOutput
				z++
			}
		}
	}
}

// barrier builds a barrier biome
func (b *Biome) barrier(logger *VerboseLogger) {
	logger.Log("Building barrier biome\n")
	logger.Log("barrier: whole\n")
	barrier1 := BARRIERS[b.rng.Intn(len(BARRIERS))]
	logger.Log("barrier: %s\n", barrier1)
	b.replaceSection(b.sectionCanvas, fmt.Sprintf("%d", b.sectNum), barrier1)
}

// Build builds the biome
func (b *Biome) Build(logger *VerboseLogger) {
	logger.Log("section #: %d\n", b.sectNum)

	if b.sectNum%2 == 0 {
		logger.Log("biome type: BARRIER\n")
		b.barrier(logger)
	} else {
		biomeType := b.rng.Intn(len(b.availableTypes))
		b.availableTypes[biomeType](logger)
	}
}

func (b *Biome) Replace() {
	for i, row := range b.grid.canvasStrings {
		for j, val := range row {
			if val == fmt.Sprintf("%d", b.sectNum) {
				b.grid.canvasStrings[i][j] = b.sectionCanvas[i][j]
			}
		}
	}
}

// randomSample selects n random items from a slice without replacement using the provided rng
func randomSample(items []string, n int, rng *rand.Rand) []string {
	if n > len(items) {
		n = len(items)
	}
	result := make([]string, n)
	perm := rng.Perm(len(items))
	for i := range n {
		result[i] = items[perm[i]]
	}
	return result
}
