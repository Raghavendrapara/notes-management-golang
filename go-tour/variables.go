package main

import (
	"fmt"
	"math/cmplx"
	"runtime"
	"time"
)

var d, e bool

func main() {
	a := 10
	b := 10
	fmt.Println(a, b, d, e)
	fmt.Println(add(a, b))
	fmt.Println(time.Date(2026, 2, 2, 10, 0, 0, 0, time.UTC))
	shorthand()
	variables()
	loops()
	runtimes()
}
func add(a, b int) (x int) {
	return a + b
}
func shorthand() {
	i := 10
	fmt.Println(i)
}

func variables() {
	z := cmplx.Sqrt(-5)
	max := uint64(1) << 63

	fmt.Println(z, max)

	fmt.Printf("Type: %T Value: %v\n", max, max)

}

func loops() {
	for i := 0; i < 10; i++ {
		// fmt.Println(i)
	}
	i := 0
	for i < 5 {
		i++
		// fmt.Println(i)
	}
}
func runtimes() {
	fmt.Printf("Running on ")

	switch runtime.GOOS {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("OS: %s.\n", runtime.GOOS)
	}
}
