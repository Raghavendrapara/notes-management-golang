package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	fmt.Println("Main: Starting Worker")

	worker(ctx)

	fmt.Println("Main: Program Finished")

}

func worker(ctx context.Context){
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker: Stopping because: %v\n",ctx.Err())
			return
		default:
			doWork()
		}
	}
}

func doWork(){
	fmt.Println("Worker: Doing Work....")
	time.Sleep(1* time.Second)
}