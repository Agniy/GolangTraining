package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

var wg sync.WaitGroup

func main() {
	numTasks := 5
	wg.Add(numTasks)

	for i := 0; i < numTasks; i++ {
		go worker(i)
	}

	go progressTracker(numTasks)

	wg.Wait()
	fmt.Println("All tasks completed.")
}

func worker(id int) {
	defer wg.Done()

	taskDuration := time.Duration(rand.Intn(5)+1) * time.Second
	fmt.Printf("Worker %d started. Estimated time: %v\n", id, taskDuration)

	for progress := 0; progress <= 100; progress += 10 {
		time.Sleep(taskDuration / 10)
		fmt.Printf("Worker %d: %d%% complete\n", id, progress)
	}

	fmt.Printf("Worker %d finished\n", id)
}

func progressTracker(numTasks int) {
	start := time.Now()
	for {
		time.Sleep(500 * time.Millisecond)
		elapsed := time.Since(start)
		remainingTasks := int(atomic.LoadInt64((*int64)(&wg.Counter)))
		completedTasks := numTasks - remainingTasks

		if completedTasks == numTasks {
			fmt.Printf("Progress: 100%% | Tasks: %d/%d | Time: %v\n", completedTasks, numTasks, elapsed.Round(time.Millisecond))
			return
		}

		progress := float64(completedTasks) / float64(numTasks) * 100
		fmt.Printf("Progress: %.1f%% | Tasks: %d/%d | Time: %v\n", progress, completedTasks, numTasks, elapsed.Round(time.Millisecond))
	}
}

// write test func
func TestProgressTracker(t *testing.T) {
	numTasks := 5
	wg.Add(numTasks)

	for i := 0; i < numTasks; i++ {
		go worker(i)
	}

	go progressTracker(numTasks)
	wg.Wait()
}
