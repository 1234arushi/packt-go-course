package calcassignment

import (
	"sync"
	"testing"
)

func TestAdd(t *testing.T) {
	numGoroutines := 10
	var (
		wg      sync.WaitGroup
		counter int
		mu      sync.Mutex
	)
	randomSlice := GenerateRandonSlice(10000000, 1, 100)
	sum := 0
	for _, value := range randomSlice {
		sum += value

	}
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go Add(randomSlice, &counter, &mu, &wg, i, numGoroutines)
	}
	wg.Wait()

	if sum != counter {
		t.Errorf("counter returned %d,expected %d", counter, sum)
	}

}
