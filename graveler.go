// Completion Time for 1000000: 0 hours, 0 minutes, 27 seconds

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const rollCount = 1000000
const maxRolls = 231
const numThreads = 8 // Adjust based on the number of cores available

func main() {
	startTime := time.Now()
	rand.Seed(time.Now().UnixNano()) // Seed random number generator
	var wg sync.WaitGroup
	var highest int
	var mu sync.Mutex // Mutex to protect access to `highest`

	// Function to perform a batch of rolls
	rollBatch := func(start int, end int) {
		defer wg.Done()
		localHighest := 0
		for i := start; i < end; i++ {
			counts := [4]int{0, 0, 0, 0}
			for k := 0; k < maxRolls; k++ {
				counts[rand.Intn(4)]++
			}
			for _, count := range counts {
				if count > localHighest {
					localHighest = count
				}
			}
		}
		mu.Lock()
		if localHighest > highest {
			highest = localHighest
		}
		mu.Unlock()
	}

	// Launch goroutines to perform rolls concurrently
	batchSize := rollCount / numThreads
	for i := 0; i < numThreads; i++ {
		start := i * batchSize
		end := start + batchSize
		if i == numThreads-1 { // Last goroutine handles any remaining work
			end = rollCount
		}
		wg.Add(1)
		go rollBatch(start, end)
	}

	wg.Wait()	
	duration := time.Now().Sub(startTime)
	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	seconds := int(duration.Seconds()) % 60
	fmt.Printf("Duration: %d hours, %d minutes, %d seconds\n", hours, minutes, seconds)
	fmt.Printf("Highest count of roll: %d\n", highest)
	fmt.Printf("Roll Count: %d\n", rollCount)
}

