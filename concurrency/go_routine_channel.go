package concurrency

import (
	"fmt"
	"strconv"
	"sync"
)

func ChannelImpl() {
	var wg sync.WaitGroup
	numRoutines := 5
	results := make(chan *ObjectSchema) // Create a channel with buffer size equal to numRoutines

	// Launch multiple goroutines
	for i := 0; i < numRoutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			object := ChannelGoRoutine(id)
			results <- object // Send the object to the results channel
		}(i)
	}

	go func() {
		fmt.Println("All goroutines completed")
		wg.Wait()      // Wait for all goroutines to finish
		close(results) // Close the channel after all goroutines are done sending data
	}()

	fmt.Printf("Length of channel: %d\n", len(results))
	for obj := range results { // Read from the channel until it's closed
		if obj != nil {
			fmt.Printf("ID: %d, Map len: %d\n", obj.ID, len(obj.PtrObj))
		}
	}
}

func ChannelGoRoutine(id int) *ObjectSchema {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in safeGoRoutine %d: %s\n", id, r)
		}
	}()

	// Function that may cause a panic
	mayPanic(id)
	fmt.Printf("safeGoRoutine %d completed normally\n", id)

	// Create a new ObjectSchema
	m := make(map[int]string)
	m[id] = strconv.FormatInt(int64(id), 10)
	o := &ObjectSchema{
		ID:     id,
		PtrObj: m,
	}
	return o
}
