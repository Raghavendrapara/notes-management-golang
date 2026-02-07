package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

func main() {
	const tasks = 1_000_000
	const maxConcurrency = 100

	var wg sync.WaitGroup
	sem := make(chan struct{}, maxConcurrency)

	for i := 1; i <= tasks; i++ {
		wg.Add(1)

		sem <- struct{}{}

		go func(n float64) {
			defer wg.Done()
			defer func() { <-sem }()

			calculate(n)
		}(float64(i))
	}

	wg.Wait()


	ch1 := make(chan string)
    ch2 := make(chan string)

    go func() {
        time.Sleep(2 * time.Second)
        ch1 <- "Data from Service A"
    }()

    go func() {
        time.Sleep(1 * time.Second)
        ch2 <- "Data from Service B"
    }()

    // select picks the FIRST channel that is ready
    select {
    case res := <-ch1:
        fmt.Println("Received:", res)
    case res := <-ch2:
        fmt.Println("Received:", res)
    case <-time.After(3 * time.Second):
        fmt.Println("Error: Operation timed out!")
    }
	quit:= make(chan struct{})
	go worker(quit)
	signalQuitAfterWork(quit)
}
func signalQuitAfterWork(stopSignal chan<- struct{}){
	time.Sleep(5*time.Second)
	close(stopSignal)
}

func calculate(n float64) {
	_ = math.Sqrt(n)
}

func worker(stopCh <-chan struct{}) {
    for {
        select {
        case <-stopCh:
            fmt.Println("Cleaning up and exiting...")
            return
        default:
            // This 'default' makes the select non-blocking.
            // It does work if no signal is present.
            doWork()
        }
    }
}
func doWork(){
	time.Sleep(1*time.Second)
}