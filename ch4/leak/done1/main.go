package main

import (
	"fmt"
	"time"
)

func main() {
	dowork := func(
		done <-chan interface{},
		strings <-chan string,
	) <-chan interface{} {
		terminated := make(chan interface{})
		go func() {
			defer fmt.Println("dowork exit.")
			defer close(terminated)
			for {
				select {
				case s := <-strings:
					// do somethings
					fmt.Println(s)
				case <-done:
					return
				}
			}
		}()
		return terminated
	}

	done := make(chan interface{})
	terminated := dowork(done, nil)

	go func() {
		// cancel later 1 min
		time.Sleep(1 * time.Second)
		fmt.Println("Canceling dowork gorouine...")
		close(done)
	}()

	<-terminated
	fmt.Println("Done.")
}

/*
	# memo

	## dowork fmt
	func(done <-chan interface{}, strings <-chan string){ do }(<-chan interface{})

	## done
	### producer
	* make done channle
	* take done another goroutine
	* close done later do something

	### consumer
	* waiting for receiving done channel
	* if receive done, return xxx
*/
