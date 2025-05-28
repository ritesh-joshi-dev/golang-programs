package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	const numTasks = 10
	var wg sync.WaitGroup
	waitCh := make(chan struct{})

	wg.Add(numTasks)

	// Launch a goroutine to perform tasks
	go func() {
		for i := 1; i <= numTasks; i++ {
			taskID := i
			go func(id int) {
				defer wg.Done()
				fmt.Printf("Task %d started\n", id)

				// Simulate work (Uncomment to trigger timeout)
				// time.Sleep(150 * time.Millisecond)

				fmt.Printf("Task %d completed\n", id)
			}(taskID)
		}

		wg.Wait()
		close(waitCh)
	}()

	// Wait for either all tasks to complete or timeout
	select {
	case <-waitCh:
		fmt.Println("✅ All tasks finished successfully.")
	case <-time.After(100 * time.Millisecond):
		fmt.Println("⏰ Timeout: Tasks took too long.")
	}
}
