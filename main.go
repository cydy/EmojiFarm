//go:build !js || !wasm

package main

import (
	"flag"
	"time"
)

func main() {
	// Define command-line flags
	seedFlag := flag.String("seed", "", "(Optional) Seed for deterministic generation (string).")
	var verbose bool
	flag.BoolVar(&verbose, "v", false, "Enable verbose output")
	flag.BoolVar(&verbose, "verbose", false, "Enable verbose output (same as -v)")
	flag.Parse()

	var seed int64
	logger := NewVerboseLogger(verbose)

	// Check for positional argument first
	args := flag.Args()
	if len(args) > 0 && *seedFlag == "" {
		// Use positional argument if no -seed flag was provided
		seed = int64(hashString(args[0]))
		logger.Log("Using seed from positional argument '%s': %d\n", args[0], seed)
	} else if *seedFlag != "" {
		// Use -seed flag if provided
		seed = int64(hashString(*seedFlag))
		logger.Log("Using seed from -seed flag '%s': %d\n", *seedFlag, seed)
	} else {
		// No seed provided, use current time
		seed = time.Now().UnixNano()
		logger.Log("Using time-based seed: %d\n", seed)
	}

	// Create a new grid with the determined seed
	grid := NewEmojiGrid(seed)

	// Build the farm
	grid.BuildFarm(logger)
}
