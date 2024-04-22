package concurrency

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type ObjectSchema struct {
	ID     int
	PtrObj map[int]string
}

func GoRoutineScript() {
	var wg sync.WaitGroup
	numRoutines := 5
	m := make([]*ObjectSchema, 0)
	var mutex sync.Mutex // Mutex to protect slice access

	// Launch multiple goroutines
	for i := 0; i < numRoutines; i++ {
		wg.Add(1) // Increment the WaitGroup counter
		go func(id int) {
			defer wg.Done() // Decrement the counter when the goroutine completes
			safeGoRoutine(id, &m, &mutex)
		}(i)
	}

	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("All goroutines completed")

	for _, obj := range m {
		fmt.Printf("ID: %d, Map len: %d\n", obj.ID, len(obj.PtrObj))
	}
}

func safeGoRoutine(id int, sharedSlice *[]*ObjectSchema, mutex *sync.Mutex) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in safeGoRoutine %d: %s\n", id, r)
		}
	}()

	// Function that may cause a panic
	mayPanic(id)
	fmt.Printf("safeGoRoutine %d completed normally\n", id)

	// Create a object and append it into the sharedSlice
	m := make(map[int]string)
	m[id] = strconv.FormatInt(int64(id), 10)
	o := &ObjectSchema{
		ID:     id,
		PtrObj: m,
	}

	mutex.Lock() // Lock before accessing shared data
	*sharedSlice = append(*sharedSlice, o)
	mutex.Unlock() // Unlock after modifying shared data
}

func mayPanic(id int) {
	// Simulate a panic situation based on an arbitrary condition
	if id%2 == 0 { // Let's say we panic on even numbers for demonstration
		time.Sleep(1 * time.Second)
		panic(fmt.Sprintf("something went wrong in goroutine %d", id))
	}
}
