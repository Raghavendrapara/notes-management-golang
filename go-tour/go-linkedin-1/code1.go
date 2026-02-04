package main

import (
	"fmt"
	"time"
)

/*
Very important to note that the main is a coroutine and ends without caring for others
so a sleep becomes eseential to ensure that the various couroutines under it get completed

But it is not recommended for production :)

*/
func main() {
	go hello()
	go end()
	//Without below it might not print anyhting
	time.Sleep(1 * time.Second)

}
func hello() {
	fmt.Println("Hello")
}
func end() {
	fmt.Println("Bye")
}
