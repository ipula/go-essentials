package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)
	doneChan <- true
	// close(doneChan)
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	doneChan <- true
	close(doneChan)
}

func main() {
	done := make(chan bool)
	// dones := make([]chan bool, 4)
	// for i := range dones {
	// 	dones[i] = make(chan bool)
	// }
	go greet("Nice to meet you!", done)
	go greet("How are you?", done)
	go slowGreet("How ... are ... you ...?", done)
	go greet("I hope you're liking the course!", done)
	// for _, done := range dones {
	// 	fmt.Println(<-done)
	// }
	for range done {
		// fmt.Println(doneChan)
	}
}
