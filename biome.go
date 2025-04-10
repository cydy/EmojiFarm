package main

import (
	"fmt"
	"math"
	"math/rand"
)

type Biome struct {
	grid           *EmojiGrid
	sectionCanvas  [][]string
	sectNum        int
	cords          [][]int
	availableTypes map[int]func()
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

	b.availableTypes = map[int]func(){
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

func (b *Biome) crops() {
	//if rand.Intn(3) == 1 {
	if true {
		fmt.Println("crops: whole")
		crop1 := PLANTS_CROPS[b.rng.Intn(len(PLANTS_CROPS))]
		fmt.Println("crop:", crop1)

		b.replaceSection(b.sectionCanvas, fmt.Sprint(b.sectNum), crop1)
	}
}

func (b *Biome) animals() {
	placementCords := b.cords
	numberOfCords := len(placementCords)

	// Fill the entire section with grass background
	for _, cord := range placementCords {
		b.sectionCanvas[cord[0]][cord[1]] = "🌱"
	}

	fillPercentage := b.rng.Float64()*(0.6-0.2) + 0.2
	numCordsToFill := int(math.Round(float64(numberOfCords) * fillPercentage))

	fmt.Printf("filling %d, %.2f, %d\n", numberOfCords, fillPercentage, numCordsToFill)

	// Shuffle the placement coordinates using b.rng
	b.rng.Shuffle(len(placementCords), func(i, j int) {
		placementCords[i], placementCords[j] = placementCords[j], placementCords[i]
	})

	var animalSet []string
	if b.rng.Intn(2) == 1 {
		fmt.Println("animals: livestock")
		animalSet = ANIMALS_FARM
	} else {
		fmt.Println("animals: birds")
		animalSet = ANIMALS_BIRDS
	}

	if b.rng.Intn(4) == 0 { // 1/4 chance one animal biome
		fmt.Println("animals #: 1")
		animal1 := animalSet[b.rng.Intn(len(animalSet))]
		fmt.Println("animal1:", animal1)
		for i := range numCordsToFill {
			cord := placementCords[i]
			b.sectionCanvas[cord[0]][cord[1]] = animal1
		}
	} else {
		animalMax := min(5, numCordsToFill)
		animalNum := b.rng.Intn(animalMax) + 1
		fmt.Println("animals #:", animalNum)

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

func (b *Biome) field() {
	placementCords := b.cords
	numberOfCords := len(placementCords)

	// Fill the entire section with grass background
	for _, cord := range placementCords {
		b.sectionCanvas[cord[0]][cord[1]] = "🌱"
	}

	fillPercentage := b.rng.Float64()*(0.8-0.2) + 0.2
	numCordsToFill := int(math.Round(float64(numberOfCords) * fillPercentage))

	fmt.Printf("filling %d, %.2f, %d\n", numberOfCords, fillPercentage, numCordsToFill)

	// Shuffle the placement coordinates using b.rng
	b.rng.Shuffle(len(placementCords), func(i, j int) {
		placementCords[i], placementCords[j] = placementCords[j], placementCords[i]
	})

	var plantSet []string
	if b.rng.Intn(2) == 1 {
		fmt.Println("plants: field")
		plantSet = PLANTS_FIELD
	} else {
		fmt.Println("plants: flowers")
		plantSet = PLANTS_FLOWERS
	}

	if b.rng.Intn(4) == 0 { // 1/4 chance one plant biome
		fmt.Println("plants #: 1")
		plant1 := plantSet[b.rng.Intn(len(plantSet))]
		fmt.Println("plant1:", plant1)
		for i := range numCordsToFill {
			cord := placementCords[i]
			b.sectionCanvas[cord[0]][cord[1]] = plant1
		}
	} else {
		plantMax := min(5, numCordsToFill)
		plantNum := b.rng.Intn(plantMax) + 1
		fmt.Println("plants #:", plantNum)

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

func (b *Biome) pond() {
	placementCords := b.cords
	numberOfCords := len(placementCords)
	fillPercentage := b.rng.Float64()*0.4 + 0.1
	numCordsToFill := int(math.Round(float64(numberOfCords) * fillPercentage))
	fmt.Printf("filling %d, %.2f, %d\n", numberOfCords, fillPercentage, numCordsToFill)

	// Shuffle placementCords using b.rng
	b.rng.Shuffle(len(placementCords), func(i, j int) {
		placementCords[i], placementCords[j] = placementCords[j], placementCords[i]
	})
	placementCords = placementCords[:numCordsToFill]

	b.replaceSection(b.sectionCanvas, fmt.Sprintf("%d", b.sectNum), "🌊")

	pawnSet := POND

	if b.rng.Intn(4) == 0 { // 3/4 chance one pawn biome
		fmt.Println("pawns #: 1")
		pawn1 := pawnSet[b.rng.Intn(len(pawnSet))]
		fmt.Println("pawn1:", pawn1)
		for _, cord := range placementCords {
			b.sectionCanvas[cord[0]][cord[1]] = pawn1
		}
	} else {
		pawnMax := min(5, numCordsToFill)
		pawnNum := b.rng.Intn(pawnMax) + 1
		fmt.Println("pawns #:", pawnNum)

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

func (b *Biome) barrier() {
	if true { // Always true for now, can be changed to match Python logic if needed
		fmt.Println("barrier: whole")
		barrier1 := BARRIERS[b.rng.Intn(len(BARRIERS))]
		fmt.Println("barrier:", barrier1)

		b.replaceSection(b.sectionCanvas, fmt.Sprintf("%d", b.sectNum), barrier1)
	}
}

func (b *Biome) Build() {
	fmt.Printf("\nsection #: %d\n", b.sectNum)

	if b.sectNum%2 == 0 {
		fmt.Println("biome type: BARRIER")
		b.barrier()
	} else {
		biomeType := b.rng.Intn(len(b.availableTypes))
		b.availableTypes[biomeType]()
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
