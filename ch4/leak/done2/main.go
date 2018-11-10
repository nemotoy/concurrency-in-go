package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	newRandomStream := func(done <-chan interface{}) <-chan int {
		randStream := make(chan int)
		go func() {
			defer fmt.Println("newRandomStream closure exited.")
			defer close(randStream)
			for {
				select {
				case randStream <- rand.Int():
				case <-done:
					return
				}
			}
		}()

		return randStream
	}

	// producer to 'end' job.
	done := make(chan interface{})
	randStream := newRandomStream(done)
	fmt.Println("3 random ints:")
	for i := 0; i <= 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}
	// job finished.
	close(done)
	time.Sleep(1 * time.Second)
}
