package calcassignment

import (
	"math/rand"
	"sync"
	"time"
)

//seed -> a starting number that computer needs to generate sequence of random numbers.

func GenerateRandonSlice(size, min, max int) []int {
	//time converted to nanoseconds
	//r -> a random engine(if not use,then a default seed would be used and it would generate same seq of random numbers)
	r := rand.New(rand.NewSource(time.Now().UnixNano())) //since time changes every moment,hence seed would also be different in each run
	randomSlice := make([]int, size)
	for i := 0; i < size; i++ {
		//Intn(n)->generates a random number between 0 and n-1
		randomSlice[i] = r.Intn(max-min+1) + min
	}
	return randomSlice
}

func Add(numbers []int, counter *int, mu *sync.Mutex, wg *sync.WaitGroup, goroutineIndex, numberGoRoutines int) {
	defer wg.Done()

	numEl := len(numbers)
	elementsPerGoRoutine := (numEl + numberGoRoutines - 1) / numberGoRoutines
	start := goroutineIndex * elementsPerGoRoutine
	end := (goroutineIndex + 1) * elementsPerGoRoutine
	if end > numEl {
		end = numEl
	}
	mu.Lock()

	for i := start; i < end; i++ {
		*counter += numbers[i]

	}
	mu.Unlock()

}
