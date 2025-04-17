//go:build js && wasm

package main

import (
	"syscall/js"
	"time"
)

func main() {
	// Get document object
	document := js.Global().Get("document")

	// Get seed value from input field
	seedInput := document.Call("getElementById", "seed")
	seedStr := seedInput.Get("value").String()

	// Get verbose setting from checkbox
	verboseCheckbox := document.Call("getElementById", "verboseToggle")
	verbose := verboseCheckbox.Get("checked").Bool()

	var seed int64
	logger := NewVerboseLogger(verbose)

	if seedStr != "" {
		seed = int64(hashString(seedStr))
		logger.Log("Using seed from input '%s': %d\n", seedStr, seed)
	} else {
		seed = time.Now().UnixNano()
		logger.Log("Using time-based seed: %d\n", seed)
	}

	// Create a new grid with the determined seed
	grid := NewEmojiGrid(seed)

	// Build the farm
	grid.BuildFarm(logger)

	// Get the grid as a string
	gridString := grid.PrintFinalEmojis(false)

	emojiGridDiv := document.Call("getElementById", "emojiGrid")

	// Set the textContent of the div
	emojiGridDiv.Set("textContent", gridString)
}
