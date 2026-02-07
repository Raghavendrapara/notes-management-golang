package main

import (
	"math"
	"sync"
)

func main() {
	tasks := 100_000_000
	workerCount := 10000

	jobs := make(chan float64, workerCount)
	var wg sync.WaitGroup

	for w := 1; w <= workerCount; w++ {
		wg.Add(1)
		go PrintRoot(jobs, &wg)
	}

	for i := 1; i <= tasks; i++ {
		jobs <- float64(i)
	}

	close(jobs)
	wg.Wait()
}

func PrintRoot(jobs <-chan float64, wg *sync.WaitGroup) {
	defer wg.Done()

	for n := range jobs {
		_ = math.Sqrt(n)
	}
}
